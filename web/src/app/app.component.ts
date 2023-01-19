import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivationEnd, NavigationExtras, Router } from '@angular/router';
import { ListData } from '@uxwb/components';
import { filter, Subscription } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit, OnDestroy {
  menuKey = 'home';

  title = '';

  menu: ListData[] = [
    { key: 'home', name: 'Home' },
    { key: 'firewall', name: 'Firewall' },
    { key: 'sites', name: 'Sites' },
    { key: 'redirect', name: 'Redirect' },
    { key: 'scheduler', name: 'Scheduler' },
  ];

  private routSub$?: Subscription;

  constructor(
    private router: Router,
  ) {
  }

  ngOnInit() {
    this.routSub$ = this.router.events
      .pipe(
        filter(e => e instanceof ActivationEnd),
      )
      .subscribe((value) => {
        if (value instanceof ActivationEnd) {
          this.menuKey = value.snapshot.routeConfig?.path || '';
          this.title = value.snapshot.routeConfig?.title?.toString() || '';
        }
      });
  }

  ngOnDestroy() {
    this.routSub$?.unsubscribe();
  }

  changeRoute(uri: string): void {
    const navigationExtras: NavigationExtras = {
      queryParamsHandling: '',
      preserveFragment: false,
    };
    this.router.navigate([uri], navigationExtras);
  }
}
