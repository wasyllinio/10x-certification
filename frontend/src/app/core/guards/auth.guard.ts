import { type CanActivateFn } from '@angular/router';

export const authGuard: CanActivateFn = () => {
  // Check if JWT token exists in localStorage
  const token = localStorage.getItem('jwt');
  
  if (!token) {
    // TODO: Implement navigation to login
    // const router = inject(Router);
    // router.navigate(['/auth/login']);
    // return false;
  }
  
  // TODO: Add token expiration check
  return true;
};

