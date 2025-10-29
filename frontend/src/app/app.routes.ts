import { Routes } from '@angular/router';
import { AppLayoutComponent } from './layout/components/app-layout/app-layout.component';

export const routes: Routes = [
  // Auth routes (without layout)
  {
    path: 'auth',
    loadChildren: () => import('./features/auth/auth.routes').then(m => m.AUTH_ROUTES)
  },
  
  // Main layout routes (with sidebar and topbar)
  {
    path: '',
    component: AppLayoutComponent,
    children: [
      {
        path: 'chargers',
        loadChildren: () => import('./features/chargers/chargers.routes').then(m => m.CHARGERS_ROUTES)
      },
      // Dashboard route (to be implemented)
      // {
      //   path: 'dashboard',
      //   loadComponent: () => import('./features/dashboard/components/dashboard/dashboard.component').then(m => m.DashboardComponent)
      // },
      // {
      //   path: 'locations',
      //   loadChildren: () => import('./features/locations/locations.routes').then(m => m.LOCATIONS_ROUTES)
      // },
      // Default redirect to chargers
      {
        path: '',
        redirectTo: 'chargers',
        pathMatch: 'full'
      }
    ]
  },
  
  // 404 route
  {
    path: '**',
    redirectTo: '/chargers'
  }
];
