import { Component, OnInit, signal, effect, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet, Router } from '@angular/router';
import { SidebarComponent } from '../sidebar/sidebar.component';
import { TopbarComponent } from '../topbar/topbar.component';
import { ConfirmDialog } from 'primeng/confirmdialog';
import { Toast } from 'primeng/toast';
import { BreadcrumbService } from '../../../core/services/breadcrumb.service';

@Component({
  selector: 'app-layout',
  standalone: true,
  imports: [CommonModule, RouterOutlet, SidebarComponent, TopbarComponent, ConfirmDialog, Toast],
  templateUrl: './app-layout.component.html',
  styleUrl: './app-layout.component.css'
})
export class AppLayoutComponent implements OnInit {
  sidebarCollapsed = signal(false);
  private router = inject(Router);
  private breadcrumbService = inject(BreadcrumbService);

  constructor() {
    // Save state to localStorage when it changes
    effect(() => {
      localStorage.setItem('sidebarCollapsed', String(this.sidebarCollapsed()));
    });
  }

  ngOnInit(): void {
    // Initialize breadcrumb service with router
    this.breadcrumbService.initialize(this.router);

    // Load sidebar state from localStorage
    const savedState = localStorage.getItem('sidebarCollapsed');
    if (savedState !== null) {
      this.sidebarCollapsed.set(savedState === 'true');
    }
  }

  onToggleSidebar(): void {
    this.sidebarCollapsed.update(value => !value);
  }
}
