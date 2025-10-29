import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-empty-state',
  standalone: true,
  imports: [CommonModule, ButtonModule],
  template: `
    <div class="empty-state-container">
      @if (icon) {
        <i [class]="icon" class="empty-icon"></i>
      }
      <h3>{{ title }}</h3>
      @if (message) {
        <p>{{ message }}</p>
      }
      @if (actionLabel && actionCallback) {
        <p-button 
          [label]="actionLabel" 
          (click)="actionCallback()"
          [severity]="actionSeverity || 'primary'">
        </p-button>
      }
    </div>
  `,
  styles: `
    .empty-state-container {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      padding: 3rem;
      text-align: center;
      min-height: 300px;
    }
    
    .empty-icon {
      font-size: 4rem;
      color: var(--p-color-text-secondary);
      margin-bottom: 1rem;
    }
    
    h3 {
      margin: 1rem 0 0.5rem;
      color: var(--p-color-text);
    }
    
    p {
      margin: 0.5rem 0 1.5rem;
      color: var(--p-color-text-secondary);
    }
  `
})
export class EmptyStateComponent {
  @Input() icon?: string;
  @Input() title: string = 'No data';
  @Input() message?: string;
  @Input() actionLabel?: string;
  @Input() actionCallback?: () => void;
  @Input() actionSeverity: 'primary' | 'secondary' | 'success' | 'info' | 'danger' = 'primary';
}

