import { Directive, EventEmitter, Output, HostListener, OnDestroy } from '@angular/core';

@Directive({
  selector: '[appDebounce]',
  standalone: true
})
export class DebounceDirective implements OnDestroy {
  @Output() debounceInput = new EventEmitter<string>();
  
  private timeoutId?: number;

  @HostListener('input', ['$event'])
  onInput(event: Event): void {
    const value = (event.target as HTMLInputElement).value;

    if (this.timeoutId) {
      clearTimeout(this.timeoutId);
    }

    this.timeoutId = window.setTimeout(() => {
      this.debounceInput.emit(value);
    }, 300);
  }

  ngOnDestroy(): void {
    if (this.timeoutId) {
      clearTimeout(this.timeoutId);
    }
  }
}

