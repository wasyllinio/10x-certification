import { Injectable, inject } from '@angular/core';
import { ConfirmationService } from 'primeng/api';

@Injectable({
  providedIn: 'root'
})
export class ConfirmService {
  private confirmationService = inject(ConfirmationService);

  confirm(header: string, message: string): Promise<boolean> {
    return new Promise((resolve) => {
      this.confirmationService.confirm({
        header,
        message,
        acceptLabel: 'Yes',
        rejectLabel: 'No',
        accept: () => resolve(true),
        reject: () => resolve(false)
      });
    });
  }
}

