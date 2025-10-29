import { Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';

export const AUTH_ROUTES: Routes = [
  {
    path: 'login',
    component: LoginComponent
  },
  // Default redirect
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
  }
];

