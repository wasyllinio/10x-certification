import { inject } from '@angular/core';
import { type CanDeactivateFn } from '@angular/router';
import { ConfirmService } from '../../shared/components/confirm-dialog/confirm.service';

export interface CanComponentDeactivate {
  canDeactivate(): boolean;
}

export const canDeactivateGuard: CanDeactivateFn<CanComponentDeactivate> = (component) => {
  if (!component) {
    return true;
  }

  // If component is clean, allow navigation
  if (component.canDeactivate()) {
    return true;
  }

  // Otherwise, show confirmation dialog
  const confirmService = inject(ConfirmService);
  return confirmService.confirm(
    'Unsaved changes',
    'You have unsaved changes. Are you sure you want to leave?'
  );
};

