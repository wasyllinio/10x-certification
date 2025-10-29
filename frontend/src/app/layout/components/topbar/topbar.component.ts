import { Component, computed, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { BreadcrumbModule } from 'primeng/breadcrumb';
import { ButtonModule } from 'primeng/button';
import { MenuModule } from 'primeng/menu';
import { MenuItem } from 'primeng/api';
import { BreadcrumbService } from '../../../core/services/breadcrumb.service';

@Component({
  selector: 'app-topbar',
  standalone: true,
  imports: [CommonModule, RouterModule, BreadcrumbModule, ButtonModule, MenuModule],
  templateUrl: './topbar.component.html',
  styleUrl: './topbar.component.css'
})
export class TopbarComponent {
  private breadcrumbService = inject(BreadcrumbService);

  breadcrumbItems = computed(() => {
    return this.breadcrumbService.breadcrumbs().map(breadcrumb => ({
      label: breadcrumb.label,
      routerLink: breadcrumb.url
    }));
  });

  userMenuItems: MenuItem[] = [
    {
      label: 'Profile',
      icon: 'pi pi-user',
      routerLink: '/profile'
    },
    {
      separator: true
    },
    {
      label: 'Logout',
      icon: 'pi pi-sign-out',
      command: () => this.handleLogout()
    }
  ];

  handleLogout(): void {
    localStorage.removeItem('jwt');
    window.location.href = '/auth/login';
  }
}

