import { Component, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';

@Component({
  selector: 'app-version-conflict-modal',
  standalone: true,
  imports: [CommonModule, ButtonModule, DialogModule],
  template: `
    <p-dialog 
      [(visible)]="visible" 
      [modal]="true" 
      [closable]="false"
      [draggable]="false"
      styleClass="version-conflict-dialog"
      [blockScroll]="true">
      <div class="version-conflict-content">
        <h3>Version Conflict</h3>
        <p>
          The resource you are trying to update has been modified by another user. 
          Please refresh the data to get the latest version.
        </p>
        <div class="actions">
          <p-button 
            label="Refresh Data" 
            (onClick)="handleRefresh()"
            severity="primary">
          </p-button>
          <p-button 
            label="Cancel" 
            (onClick)="handleCancel()"
            severity="secondary">
          </p-button>
        </div>
      </div>
    </p-dialog>
  `,
  styles: `
    .version-conflict-content {
      padding: 1rem;
    }
    
    h3 {
      margin-top: 0;
      color: var(--p-color-text);
    }
    
    p {
      margin-bottom: 1.5rem;
      color: var(--p-color-text-secondary);
    }
    
    .actions {
      display: flex;
      gap: 0.5rem;
      justify-content: flex-end;
    }
  `
})
export class VersionConflictModalComponent {
  visible = signal(false);
  refreshCallback?: () => void;
  cancelCallback?: () => void;

  show(refreshCallback?: () => void, cancelCallback?: () => void): void {
    this.refreshCallback = refreshCallback;
    this.cancelCallback = cancelCallback;
    this.visible.set(true);
  }

  handleRefresh(): void {
    this.visible.set(false);
    if (this.refreshCallback) {
      this.refreshCallback();
    }
  }

  handleCancel(): void {
    this.visible.set(false);
    if (this.cancelCallback) {
      this.cancelCallback();
    }
  }
}

