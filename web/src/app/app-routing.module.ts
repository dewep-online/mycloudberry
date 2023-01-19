import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ErrorPageComponent } from './modules/error-page/error-page.component';
import { FirewallComponent } from './modules/firewall/firewall.component';
import { HomeComponent } from './modules/home/home.component';
import { RedirectComponent } from './modules/redirect/redirect.component';
import { SchedulerComponent } from './modules/scheduler/scheduler.component';
import { SitesComponent } from './modules/sites/sites.component';

const routes: Routes = [
  { path: 'home', component: HomeComponent, title: 'Home' },
  { path: 'firewall', component: FirewallComponent, title: 'Firewall' },
  { path: 'redirect', component: RedirectComponent, title: 'Redirect' },
  { path: 'sites', component: SitesComponent, title: 'Sites' },
  { path: 'scheduler', component: SchedulerComponent, title: 'Scheduler' },
  //---
  { path: '', pathMatch: 'full', redirectTo: '/home' },
  { path: '**', component: ErrorPageComponent, title: 'Errors' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule],
})
export class AppRoutingModule {
}
