import { Directive, Input, OnInit, TemplateRef, ViewContainerRef, inject, effect } from '@angular/core';
import { RoleService } from '../services/role.service';

@Directive({
  selector: '[appRole]',
  standalone: true
})
export class RoleDirective implements OnInit {
  @Input({ required: true }) appRole!: string[] | string;
  
  private roleService = inject(RoleService);
  private templateRef = inject(TemplateRef<unknown>);
  private viewContainer = inject(ViewContainerRef);

  ngOnInit(): void {
    effect(() => {
      const userRole = this.roleService.currentRole();
      const allowedRoles = Array.isArray(this.appRole) ? this.appRole : [this.appRole];
      
      if (userRole && allowedRoles.includes(userRole)) {
        this.viewContainer.createEmbeddedView(this.templateRef);
      } else {
        this.viewContainer.clear();
      }
    });
  }
}

