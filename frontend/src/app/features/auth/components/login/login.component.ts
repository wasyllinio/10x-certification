import { Component, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { PasswordModule } from 'primeng/password';
import { MessageModule } from 'primeng/message';
import { MessageService } from 'primeng/api';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../../../environments/environment';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    InputTextModule,
    ButtonModule,
    CardModule,
    PasswordModule,
    MessageModule
  ],
  template: `
    <div class="login-container">
      <p-card>
        <div class="login-content">
          <h1 class="login-title">Logowanie</h1>
          <p class="login-subtitle">Zaloguj się do systemu 10x Certification</p>
          
          <form [formGroup]="loginForm" (ngSubmit)="onSubmit()">
            <div class="field-wrapper">
              <label for="email">Email</label>
              <input 
                id="email"
                type="email" 
                pInputText 
                formControlName="email"
                placeholder="twoj@email.pl"
                autocomplete="email"
                [attr.aria-required]="true"
                aria-describedby="email-error"
              />
              @if (loginForm.get('email')?.invalid && loginForm.get('email')?.touched) {
                <small id="email-error" class="p-error" role="alert">Email jest wymagany</small>
              }
            </div>
            
            <div class="field-wrapper">
              <label for="password">Hasło</label>
              <p-password 
                id="password"
                formControlName="password"
                placeholder="Hasło"
                [feedback]="false"
                [toggleMask]="true"
                autocomplete="current-password"
                [attr.aria-required]="true"
                inputId="password-input"
              />
              @if (loginForm.get('password')?.invalid && loginForm.get('password')?.touched) {
                <small id="password-error" class="p-error" role="alert">Hasło jest wymagane</small>
              }
            </div>
            
            @if (error()) {
              <p-message severity="error" [text]="error()!"></p-message>
            }
            
            <div class="buttons-wrapper">
              <p-button 
                type="submit"
                label="Zaloguj się" 
                [disabled]="loginForm.invalid || loading()"
                [loading]="loading()"
                styleClass="w-full">
              </p-button>
              
              <p-button 
                type="button"
                label="Pomiń logowanie (dev)" 
                severity="secondary"
                styleClass="w-full"
                (click)="skipLogin()">
              </p-button>
            </div>
          </form>
        </div>
      </p-card>
    </div>
  `,
  styles: `
    .login-container {
      display: flex;
      align-items: center;
      justify-content: center;
      min-height: 100vh;
      background: var(--p-color-surface-section);
      padding: 1rem;
    }
    
    :host ::ng-deep .p-card {
      width: 100%;
      max-width: 400px;
    }
    
    .login-content {
      padding: 1.5rem;
    }
    
    .login-title {
      font-size: 2rem;
      font-weight: 600;
      margin: 0 0 0.5rem 0;
      color: var(--p-color-text);
    }
    
    .login-subtitle {
      color: var(--p-color-text-secondary);
      margin: 0 0 2rem 0;
    }
    
    .field-wrapper {
      margin-bottom: 1.5rem;
    }
    
    label {
      display: block;
      margin-bottom: 0.5rem;
      color: var(--p-color-text);
      font-weight: 500;
    }
    
    .buttons-wrapper {
      display: flex;
      flex-direction: column;
      gap: 0.5rem;
      margin-top: 2rem;
    }
  `
})
export class LoginComponent {
  private fb = inject(FormBuilder);
  private router = inject(Router);
  private http = inject(HttpClient);
  private messageService = inject(MessageService);
  
  loginForm: FormGroup = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required]]
  });
  
  loading = signal(false);
  error = signal<string | null>(null);
  
  constructor() {
    // Check if already logged in
    const token = localStorage.getItem('jwt');
    if (token) {
      this.router.navigate(['/chargers']);
    }
  }
  
  onSubmit(): void {
    if (this.loginForm.invalid) {
      this.markFormGroupTouched(this.loginForm);
      return;
    }
    
    this.loading.set(true);
    this.error.set(null);
    
    const { email, password } = this.loginForm.value;
    
    this.http.post<{ token: string }>(`${environment.apiUrl}/auth/login`, { email, password })
      .subscribe({
        next: (response) => {
          localStorage.setItem('jwt', response.token);
          this.loading.set(false);
          this.router.navigate(['/chargers']);
        },
        error: (err) => {
          this.error.set(err.error?.message || 'Nieprawidłowy email lub hasło');
          this.loading.set(false);
        }
      });
  }
  
  skipLogin(): void {
    // For development: skip authentication
    localStorage.setItem('jwt', 'dev-token');
    this.router.navigate(['/chargers']);
  }
  
  private markFormGroupTouched(formGroup: FormGroup): void {
    Object.keys(formGroup.controls).forEach(key => {
      const control = formGroup.get(key);
      control?.markAsTouched();
    });
  }
}

