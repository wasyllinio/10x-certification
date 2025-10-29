import { Injectable, signal, computed } from '@angular/core';
import { Router, NavigationEnd, ActivatedRoute } from '@angular/router';
import { filter, map } from 'rxjs/operators';

export interface BreadcrumbItem {
  label: string;
  url: string;
}

@Injectable({
  providedIn: 'root'
})
export class BreadcrumbService {
  private router = signal<Router | null>(null);
  private items = signal<BreadcrumbItem[]>([]);

  public readonly breadcrumbs = computed(() => this.items());

  initialize(router: Router): void {
    this.router.set(router);

    router.events
      .pipe(
        filter(event => event instanceof NavigationEnd),
        map(() => router.routerState.root),
        map((route: ActivatedRoute) => {
          while (route.firstChild) {
            route = route.firstChild;
          }
          return route;
        })
      )
      .subscribe((route: ActivatedRoute) => {
        this.buildBreadcrumb(route);
      });
  }

  private buildBreadcrumb(route: ActivatedRoute): void {
    const breadcrumbs: BreadcrumbItem[] = [];
    
    if (route.routeConfig?.data?.['breadcrumb']) {
      breadcrumbs.push({
        label: route.snapshot.data['breadcrumb'],
        url: this.getUrl(route)
      });
    }

    this.items.set(breadcrumbs);
  }

  private getUrl(route: ActivatedRoute): string {
    const urlSegments = route.snapshot.url;
    const parent = route.parent;
    
    const segments: string[] = [];
    
    if (parent) {
      const parentUrl = this.getUrl(parent);
      segments.push(...parentUrl.split('/').filter(s => s));
    }
    
    urlSegments.forEach(segment => {
      segments.push(segment.path);
    });
    
    return '/' + segments.join('/');
  }
}

