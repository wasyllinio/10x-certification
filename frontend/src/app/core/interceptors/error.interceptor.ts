import { HttpInterceptorFn, HttpErrorResponse } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, throwError } from 'rxjs';
import { MessageService } from 'primeng/api';
import { VersionConflictService } from '../../shared/components/version-conflict-modal/version-conflict.service';

export const errorInterceptor: HttpInterceptorFn = (req, next) => {
  const router = inject(Router);
  const messageService = inject(MessageService);
  const versionConflictService = inject(VersionConflictService);

  return next(req).pipe(
    catchError((error: HttpErrorResponse) => {
      if (error.status === 401) {
        // Unauthorized - redirect to login
        localStorage.removeItem('jwt');
        router.navigate(['/auth/login']);
        messageService.add({
          severity: 'error',
          summary: 'Session expired',
          detail: 'Please log in again'
        });
      } else if (error.status === 403) {
        // Forbidden
        messageService.add({
          severity: 'error',
          summary: 'Access denied',
          detail: 'You do not have permission to perform this action'
        });
      } else if (error.status === 404) {
        // Not found
        messageService.add({
          severity: 'warn',
          summary: 'Not found',
          detail: 'The requested resource was not found'
        });
      } else if (error.status === 409) {
        // Version conflict - optimistic locking
        versionConflictService.showConflictModal();
      } else if (error.status >= 500) {
        // Server errors
        messageService.add({
          severity: 'error',
          summary: 'Server error',
          detail: 'An unexpected error occurred. Please try again later.'
        });
      } else if (error.status === 400) {
        // Bad request - validation errors
        const errorMessage = error.error?.message || 'Invalid request';
        messageService.add({
          severity: 'error',
          summary: 'Validation error',
          detail: errorMessage
        });
      }

      return throwError(() => error);
    })
  );
};

