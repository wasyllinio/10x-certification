import { Component } from '@angular/core';
import { ProgressSpinnerModule } from 'primeng/progressspinner';

@Component({
  selector: 'app-loading-spinner',
  standalone: true,
  imports: [ProgressSpinnerModule],
  template: `
    <div class="loading-container">
      <p-progressSpinner 
        class="spinner" 
        strokeWidth="4" 
        fill="transparent" 
        animationDuration=".5s">
      </p-progressSpinner>
    </div>
  `,
  styles: `
    .loading-container {
      display: flex;
      justify-content: center;
      align-items: center;
      padding: 2rem;
    }
    
    .spinner {
      width: 50px;
      height: 50px;
    }
  `
})
export class LoadingSpinnerComponent {}

