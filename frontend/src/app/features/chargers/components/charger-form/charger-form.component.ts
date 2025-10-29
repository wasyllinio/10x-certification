import { Component, OnInit, signal, inject, effect } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormArray, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { Router, ActivatedRoute } from '@angular/router';

import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { FieldsetModule } from 'primeng/fieldset';
import { SelectModule } from 'primeng/select';
import { InputNumberModule } from 'primeng/inputnumber';
import { MessageModule } from 'primeng/message';
import { ToastModule } from 'primeng/toast';
import { ProgressSpinnerModule } from 'primeng/progressspinner';

import { ChargersService } from '../../services/chargers.service';
import { MessageService } from 'primeng/api';
import { ChargerDetailResponse } from '../../models/charger.models';
import { ConnectorType, ConnectorStandard } from '../../../../../types';
import { CanComponentDeactivate } from '../../../../core/guards/can-deactivate.guard';

@Component({
  selector: 'app-charger-form',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    InputTextModule,
    ButtonModule,
    CardModule,
    FieldsetModule,
    SelectModule,
    InputNumberModule,
    MessageModule,
    ToastModule,
    ProgressSpinnerModule
  ],
  templateUrl: './charger-form.component.html',
  styleUrl: './charger-form.component.css'
})
export class ChargerFormComponent implements OnInit, CanComponentDeactivate {
  private fb = inject(FormBuilder);
  private chargersService = inject(ChargersService);
  private router = inject(Router);
  private route = inject(ActivatedRoute);
  private messageService = inject(MessageService);

  public isEditMode = signal<boolean>(false);
  public form!: FormGroup;
  public loading = this.chargersService.loading;
  public error = this.chargersService.error;
  public chargerId: string | null = null;

  // Dropdown options
  public connectorTypeOptions = [
    { label: 'CCS', value: 'CCS' as ConnectorType },
    { label: 'Type2', value: 'Type2' as ConnectorType },
    { label: 'Chademo', value: 'Chademo' as ConnectorType }
  ];

  public connectorStandardOptions = [
    { label: 'AC Single Phase', value: 'AC_1P' as ConnectorStandard },
    { label: 'AC Three Phase', value: 'AC_3P' as ConnectorStandard },
    { label: 'DC', value: 'DC' as ConnectorStandard }
  ];

  ngOnInit(): void {
    this.initializeForm();
    
    // Check if in edit mode
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id && params.get('path')?.includes('edit')) {
        this.isEditMode.set(true);
        this.chargerId = id;
        this.loadChargerData(id);
      }
    });
  }

  /**
   * Initialize form with empty values
   */
  initializeForm(): void {
    this.form = this.fb.group({
      vendor: ['', [Validators.required, Validators.maxLength(100)]],
      model: ['', [Validators.required, Validators.maxLength(100)]],
      serial_number: ['', [Validators.required, Validators.maxLength(50)]],
      connectors: this.fb.array([
        this.createConnectorFormGroup()
      ])
    });
  }

  /**
   * Create a connector form group
   */
  createConnectorFormGroup(): FormGroup {
    return this.fb.group({
      connector_id: [1, [Validators.required, Validators.min(1), Validators.max(999)]],
      power: [0, [Validators.required, Validators.min(0.1)]],
      voltage: [0, [Validators.required, Validators.min(1)]],
      amperage: [0, [Validators.required, Validators.min(0.1)]],
      connector_type: ['', Validators.required],
      connector_standard: ['', Validators.required]
    });
  }

  /**
   * Get connectors FormArray
   */
  get connectors(): FormArray {
    return this.form.get('connectors') as FormArray;
  }

  /**
   * Add new connector
   */
  addConnector(): void {
    const newConnectorNumber = this.connectors.length + 1;
    const connectorGroup = this.createConnectorFormGroup();
    connectorGroup.patchValue({ connector_id: newConnectorNumber });
    this.connectors.push(connectorGroup);
  }

  /**
   * Remove connector
   */
  removeConnector(index: number): void {
    if (this.connectors.length > 1) {
      this.connectors.removeAt(index);
    } else {
      this.messageService.add({
        severity: 'warn',
        summary: 'Cannot remove',
        detail: 'At least one connector is required'
      });
    }
  }

  /**
   * Load charger data for editing
   */
  loadChargerData(id: string): void {
    this.chargersService.getChargerById(id);
    
    // Watch for charger changes using effect
    effect(() => {
      const charger = this.chargersService.selectedCharger();
      if (charger) {
        this.populateForm(charger);
      }
    });
  }

  /**
   * Populate form with charger data
   */
  populateForm(charger: ChargerDetailResponse): void {
    // Populate basic fields
    this.form.patchValue({
      vendor: charger.vendor,
      model: charger.model,
      serial_number: charger.serial_number
    });

    // Clear existing connectors
    while (this.connectors.length !== 0) {
      this.connectors.removeAt(0);
    }

    // Add connectors from charger data
    if (charger.connectors && charger.connectors.length > 0) {
      charger.connectors.forEach((connector) => {
        const connectorGroup = this.createConnectorFormGroup();
        connectorGroup.patchValue(connector);
        this.connectors.push(connectorGroup);
      });
    }
  }

  /**
   * Submit form
   */
  onSubmit(): void {
    if (this.form.invalid) {
      this.markFormGroupTouched(this.form);
      this.messageService.add({
        severity: 'error',
        summary: 'Validation error',
        detail: 'Please fill in all required fields correctly'
      });
      return;
    }

    const formValue = this.form.value;

    if (this.isEditMode() && this.chargerId) {
      // TODO: Add version field for optimistic locking
      const updateData = {
        ...formValue,
        version: 1 // This should come from the loaded charger
      };
      
      this.chargersService.updateCharger(this.chargerId, updateData);
    } else {
      this.chargersService.createCharger(formValue);
    }

    // Redirect after success
    this.messageService.add({
      severity: 'success',
      summary: 'Success',
      detail: this.isEditMode() ? 'Charger updated successfully' : 'Charger created successfully'
    });

    setTimeout(() => {
      if (this.isEditMode() && this.chargerId) {
        this.router.navigate(['/chargers', this.chargerId]);
      } else {
        this.router.navigate(['/chargers']);
      }
    }, 1000);
  }

  /**
   * Mark form group as touched for validation
   */
  markFormGroupTouched(formGroup: FormGroup): void {
    Object.keys(formGroup.controls).forEach(key => {
      const control = formGroup.get(key);
      control?.markAsTouched();

      if (control instanceof FormGroup || control instanceof FormArray) {
        this.markFormGroupTouched(control as FormGroup);
      }
    });
  }

  /**
   * Cancel and go back
   */
  cancel(): void {
    if (this.isEditMode() && this.chargerId) {
      this.router.navigate(['/chargers', this.chargerId]);
    } else {
      this.router.navigate(['/chargers']);
    }
  }

  /**
   * Check if form is dirty for CanDeactivate guard
   */
  canDeactivate(): boolean {
    return !this.form.dirty;
  }

  /**
   * Get field error message
   */
  getFieldError(fieldName: string): string {
    const field = this.form.get(fieldName);
    if (field?.hasError('required') && field.touched) {
      return 'This field is required';
    }
    if (field?.hasError('maxlength') && field.touched) {
      return `Max length is ${field.errors?.['maxlength'].requiredLength} characters`;
    }
    if (field?.hasError('min') && field.touched) {
      return `Minimum value is ${field.errors?.['min'].min}`;
    }
    return '';
  }

  /**
   * Get connector field error message for FormArray
   */
  getConnectorFieldError(connectorIndex: number, fieldName: string): string {
    const connector = this.connectors.at(connectorIndex) as FormGroup;
    if (!connector) return '';
    
    const field = connector.get(fieldName);
    if (!field) return '';

    if (field.hasError('required') && field.touched) {
      return 'This field is required';
    }
    if (field.hasError('maxlength') && field.touched) {
      return `Max length is ${field.errors?.['maxlength'].requiredLength} characters`;
    }
    if (field.hasError('min') && field.touched) {
      return `Minimum value is ${field.errors?.['min'].min}`;
    }
    return '';
  }
}

