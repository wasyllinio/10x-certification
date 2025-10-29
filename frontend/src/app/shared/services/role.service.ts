import { Injectable, signal, computed } from '@angular/core';

export type UserRole = 'admin' | 'owner';

@Injectable({
  providedIn: 'root'
})
export class RoleService {
  private role = signal<UserRole | null>(null);

  public readonly currentRole = computed(() => this.role());
  public readonly isAdmin = computed(() => this.role() === 'admin');
  public readonly isOwner = computed(() => this.role() === 'owner');

  setRole(role: UserRole): void {
    this.role.set(role);
  }

  hasRole(allowedRoles: UserRole[]): boolean {
    const current = this.role();
    return current ? allowedRoles.includes(current) : false;
  }

  clearRole(): void {
    this.role.set(null);
  }

  parseRoleFromToken(token: string): UserRole | null {
    try {
      const payload = JSON.parse(atob(token.split('.')[1]));
      return payload.role || null;
    } catch {
      return null;
    }
  }
}

