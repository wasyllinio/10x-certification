import { Routes } from '@angular/router';
import { ChargersListComponent } from './components/chargers-list/chargers-list.component';
import { ChargerFormComponent } from './components/charger-form/charger-form.component';
import { ChargerDetailComponent } from './components/charger-detail/charger-detail.component';
import { authGuard } from '../../core/guards/auth.guard';
import { canDeactivateGuard } from '../../core/guards/can-deactivate.guard';

export const CHARGERS_ROUTES: Routes = [
  {
    path: '',
    component: ChargersListComponent,
    canActivate: [authGuard],
    data: { breadcrumb: 'Stacje' }
  },
  {
    path: 'new',
    component: ChargerFormComponent,
    canActivate: [authGuard],
    canDeactivate: [canDeactivateGuard],
    data: { breadcrumb: 'Nowa stacja' }
  },
  {
    path: ':id',
    component: ChargerDetailComponent,
    canActivate: [authGuard],
    data: { breadcrumb: 'Szczegóły stacji' }
  },
  {
    path: ':id/edit',
    component: ChargerFormComponent,
    canActivate: [authGuard],
    canDeactivate: [canDeactivateGuard],
    data: { breadcrumb: 'Edycja stacji' }
  }
];

