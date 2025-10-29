import { Component, OnInit, signal, effect, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';

import { TableModule } from 'primeng/table';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { SelectModule } from 'primeng/select';
import { CardModule } from 'primeng/card';
import { BadgeModule } from 'primeng/badge';
import { MessageService } from 'primeng/api';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { ToastModule } from 'primeng/toast';
import { ToolbarModule } from 'primeng/toolbar';
import { ProgressSpinnerModule } from 'primeng/progressspinner';

import { ChargersService } from '../../services/chargers.service';
import { DebounceDirective } from '../../../../shared/directives/debounce.directive';
import { ConfirmService } from '../../../../shared/components/confirm-dialog/confirm.service';
import { Status } from '../../models/charger.models';
import { ChargerListItem } from '../../models/charger.models';
import { TableLazyLoadEvent } from 'primeng/table';

@Component({
  selector: 'app-chargers-list',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    TableModule,
    InputTextModule,
    ButtonModule,
    SelectModule,
    CardModule,
    BadgeModule,
    ConfirmDialogModule,
    ToastModule,
    ToolbarModule,
    ProgressSpinnerModule,
    DebounceDirective
  ],
  templateUrl: './chargers-list.component.html',
  styleUrl: './chargers-list.component.css'
})
export class ChargersListComponent implements OnInit {
  private chargersService = inject(ChargersService);
  private router = inject(Router);
  private messageService = inject(MessageService);
  private confirmService = inject(ConfirmService);

  // Filters
  public searchText = signal<string>('');
  public statusFilter = signal<Status | null>(null);

  // Pagination
  public currentPage = signal<number>(1);
  public itemsPerPage = signal<number>(20);

  // Status options
  public statusOptions = [
    { label: 'All', value: null },
    { label: 'W magazynie', value: 'warehouse' as Status },
    { label: 'Przypisana', value: 'assigned' as Status }
  ];

  // Computed signals
  public chargers = this.chargersService.chargers;
  public loading = this.chargersService.loading;
  public pagination = this.chargersService.pagination;

  constructor() {
    // Effect to reload when filters change
    effect(() => {
      const search = this.searchText();
      const status = this.statusFilter();
      const page = this.currentPage();

      if (search !== undefined || status !== undefined || page !== undefined) {
        this.loadChargers();
      }
    });
  }

  ngOnInit(): void {
    this.loadChargers();
  }

  /**
   * Load chargers with current filters
   */
  loadChargers(): void {
    this.chargersService.getChargers({
      page: this.currentPage(),
      limit: this.itemsPerPage(),
      search: this.searchText() || undefined,
      status: this.statusFilter() || undefined
    });
  }

  /**
   * Handle search input with debounce
   */
  onSearchChange(search: string): void {
    this.searchText.set(search);
    this.currentPage.set(1); // Reset to first page
  }

  /**
   * Handle status filter change
   */
  onStatusChange(status: Status | null): void {
    this.statusFilter.set(status);
    this.currentPage.set(1); // Reset to first page
  }

  /**
   * Clear all filters
   */
  clearFilters(): void {
    this.searchText.set('');
    this.statusFilter.set(null);
    this.currentPage.set(1);
  }

  /**
   * Get active filters count
   */
  get activeFiltersCount(): number {
    let count = 0;
    if (this.searchText()) count++;
    if (this.statusFilter()) count++;
    return count;
  }

  /**
   * Get connector count for display
   */
  getConnectorCount(charger: ChargerListItem): number {
    return charger.connectors?.length || 0;
  }

  /**
   * Get status badge label
   */
  getStatusLabel(locationId: string | null): string {
    return locationId ? 'Przypisana' : 'W magazynie';
  }

  /**
   * Get status severity for badge
   */
  getStatusSeverity(locationId: string | null): string {
    return locationId ? 'success' : 'info';
  }

  /**
   * View charger details
   */
  viewCharger(id: string): void {
    this.router.navigate(['/chargers', id]);
  }

  /**
   * Edit charger
   */
  editCharger(id: string, event: Event): void {
    event.stopPropagation();
    this.router.navigate(['/chargers', id, 'edit']);
  }

  /**
   * Delete charger with confirmation
   */
  async deleteCharger(id: string, event: Event): Promise<void> {
    event.stopPropagation();

    const confirmed = await this.confirmService.confirm(
      'Usuń stację',
      'Czy na pewno chcesz usunąć tę stację? Ta akcja jest nieodwracalna.'
    );

    if (confirmed) {
      this.chargersService.deleteCharger(id);
      
      this.messageService.add({
        severity: 'success',
        summary: 'Usunięto',
        detail: 'Stacja została usunięta pomyślnie'
      });
    }
  }

  /**
   * Navigate to create new charger
   */
  navigateToNew(): void {
    this.router.navigate(['/chargers/new']);
  }

  /**
   * Handle pagination
   */
  onPageChange(event: TableLazyLoadEvent): void {
    const first = event.first ?? 0;
    const rows = event.rows ?? this.itemsPerPage();
    const page = first / rows;
    this.currentPage.set(page + 1);
    this.itemsPerPage.set(rows);
  }
}

