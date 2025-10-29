import { Injectable, signal, inject } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from '../../../../environments/environment';
import { 
  ChargerListResponse, 
  ChargerDetailResponse, 
  CreateChargerRequest, 
  UpdateChargerRequest,
  ChargerListItem 
} from '../../../../types';

@Injectable({
  providedIn: 'root'
})
export class ChargersService {
  private http = inject(HttpClient);
  private readonly baseUrl = `${environment.apiUrl}/chargers`;

  // Signals for state management
  public chargers = signal<ChargerListItem[]>([]);
  public loading = signal<boolean>(false);
  public error = signal<string | null>(null);
  
  public selectedCharger = signal<ChargerDetailResponse | null>(null);
  public pagination = signal({
    page: 1,
    limit: 20,
    total: 0,
    has_next: false
  });

  /**
   * Get list of chargers with pagination and filters
   */
  getChargers(params: {
    page?: number;
    limit?: number;
    search?: string;
    status?: 'warehouse' | 'assigned';
  } = {}): void {
    this.loading.set(true);
    this.error.set(null);

    let httpParams = new HttpParams()
      .set('page', params.page?.toString() || '1')
      .set('limit', params.limit?.toString() || '20');

    if (params.search) {
      httpParams = httpParams.set('search', params.search);
    }
    if (params.status) {
      httpParams = httpParams.set('status', params.status);
    }

    this.http.get<ChargerListResponse>(this.baseUrl, { params: httpParams })
      .subscribe({
        next: (response) => {
          this.chargers.set(response.data);
          this.pagination.set(response.pagination);
          this.loading.set(false);
        },
        error: (err) => {
          this.error.set(err.message || 'Failed to load chargers');
          this.loading.set(false);
        }
      });
  }

  /**
   * Get single charger by ID
   */
  getChargerById(id: string): void {
    this.loading.set(true);
    this.error.set(null);

    this.http.get<ChargerDetailResponse>(`${this.baseUrl}/${id}`)
      .subscribe({
        next: (charger) => {
          this.selectedCharger.set(charger);
          this.loading.set(false);
        },
        error: (err) => {
          this.error.set(err.message || 'Failed to load charger');
          this.loading.set(false);
        }
      });
  }

  /**
   * Create new charger
   */
  createCharger(data: CreateChargerRequest): void {
    this.loading.set(true);
    this.error.set(null);

    this.http.post<{ id: string }>(this.baseUrl, data)
      .subscribe({
        next: () => {
          this.loading.set(false);
          // Refresh list after creation
          this.getChargers();
        },
        error: (err) => {
          this.error.set(err.message || 'Failed to create charger');
          this.loading.set(false);
        }
      });
  }

  /**
   * Update existing charger
   */
  updateCharger(id: string, data: UpdateChargerRequest): void {
    this.loading.set(true);
    this.error.set(null);

    this.http.put<{ id: string }>(`${this.baseUrl}/${id}`, data)
      .subscribe({
        next: () => {
          this.loading.set(false);
          // Refresh current charger details
          this.getChargerById(id);
        },
        error: (err) => {
          this.error.set(err.message || 'Failed to update charger');
          this.loading.set(false);
        }
      });
  }

  /**
   * Delete charger
   */
  deleteCharger(id: string): void {
    this.loading.set(true);
    this.error.set(null);

    this.http.delete(`${this.baseUrl}/${id}`)
      .subscribe({
        next: () => {
          this.loading.set(false);
          // Remove from local list
          this.chargers.update(list => list.filter(c => c.id !== id));
          // Update pagination
          this.pagination.update(p => ({ ...p, total: p.total - 1 }));
        },
        error: (err) => {
          this.error.set(err.message || 'Failed to delete charger');
          this.loading.set(false);
        }
      });
  }

  /**
   * Clear selected charger
   */
  clearSelectedCharger(): void {
    this.selectedCharger.set(null);
  }
}

