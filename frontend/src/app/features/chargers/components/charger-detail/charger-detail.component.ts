import { Component, OnInit, signal, effect, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, ActivatedRoute } from '@angular/router';

import { CardModule } from 'primeng/card';
import { PanelModule } from 'primeng/panel';
import { ButtonModule } from 'primeng/button';
import { TableModule } from 'primeng/table';
import { BadgeModule } from 'primeng/badge';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { ToastModule } from 'primeng/toast';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { MessageService } from 'primeng/api';

import { ChargersService } from '../../services/chargers.service';
import { ConnectorTypePipe } from '../../../../shared/pipes/connector-type.pipe';
import { ConnectorStandardPipe } from '../../../../shared/pipes/connector-standard.pipe';
import { ConfirmService } from '../../../../shared/components/confirm-dialog/confirm.service';
import { ChargerDetailResponse } from '../../models/charger.models';

@Component({
  selector: 'app-charger-detail',
  standalone: true,
  imports: [
    CommonModule,
    CardModule,
    PanelModule,
    ButtonModule,
    TableModule,
    BadgeModule,
    ConfirmDialogModule,
    ToastModule,
    ProgressSpinnerModule,
    ConnectorTypePipe,
    ConnectorStandardPipe
  ],
  templateUrl: './charger-detail.component.html',
  styleUrl: './charger-detail.component.css'
})
export class ChargerDetailComponent implements OnInit {
  private chargersService = inject(ChargersService);
  private router = inject(Router);
  private route = inject(ActivatedRoute);
  private messageService = inject(MessageService);
  private confirmService = inject(ConfirmService);

  public chargerId: string | null = null;
  public charger = signal<ChargerDetailResponse | null>(null);
  public loading = this.chargersService.loading;
  public expandableRows = signal<boolean[]>([]);

  constructor() {
    // Update charger signal when service updates it
    effect(() => {
      const selectedCharger = this.chargersService.selectedCharger();
      if (selectedCharger) {
        this.charger.set(selectedCharger);
      }
    });
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.chargerId = id;
        this.chargersService.getChargerById(id);
      }
    });
  }

  /**
   * Toggle connector row expansion
   */
  toggleRow(index: number): void {
    const expanded = [...this.expandableRows()];
    expanded[index] = !expanded[index];
    this.expandableRows.set(expanded);
  }

  /**
   * Check if row is expanded
   */
  isRowExpanded(index: number): boolean {
    return this.expandableRows()[index] || false;
  }

  /**
   * Navigate to list
   */
  goToList(): void {
    this.router.navigate(['/chargers']);
  }

  /**
   * Navigate to edit mode
   */
  editCharger(): void {
    if (this.chargerId) {
      this.router.navigate(['/chargers', this.chargerId, 'edit']);
    }
  }

  /**
   * Delete charger with confirmation
   */
  async deleteCharger(): Promise<void> {
    const confirmed = await this.confirmService.confirm(
      'Usuń stację',
      'Czy na pewno chcesz usunąć tę stację? Ta akcja jest nieodwracalna.'
    );

    if (confirmed && this.chargerId) {
      this.chargersService.deleteCharger(this.chargerId);
      
      this.messageService.add({
        severity: 'success',
        summary: 'Usunięto',
        detail: 'Stacja została usunięta pomyślnie'
      });

      setTimeout(() => {
        this.router.navigate(['/chargers']);
      }, 1000);
    }
  }

  /**
   * Get status label
   */
  getStatusLabel(): string {
    const charger = this.charger();
    return charger?.location_id ? 'Przypisana' : 'W magazynie';
  }

  /**
   * Get status severity for badge
   */
  getStatusSeverity(): string {
    const charger = this.charger();
    return charger?.location_id ? 'success' : 'info';
  }
}

