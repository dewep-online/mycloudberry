import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { ServiceWorkerModule } from '@angular/service-worker';
import { UxwbComponentsModule } from '@uxwb/components';
import { UxwbServicesModule } from '@uxwb/services';
import { environment } from '../environments/environment';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ErrorPageComponent } from './modules/error-page/error-page.component';
import { FirewallComponent } from './modules/firewall/firewall.component';
import { HomeComponent } from './modules/home/home.component';
import { RedirectComponent } from './modules/redirect/redirect.component';
import { SchedulerComponent } from './modules/scheduler/scheduler.component';
import { SitesComponent } from './modules/sites/sites.component';

@NgModule({
  declarations: [
    AppComponent,
    FirewallComponent,
    HomeComponent,
    ErrorPageComponent,
    SitesComponent,
    RedirectComponent,
    SchedulerComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    UxwbComponentsModule,
    UxwbServicesModule.forRoot({ ajaxPrefixUrl:'/api', webSocketUrl:'/ws' }),
    ServiceWorkerModule.register('ngsw-worker.js', {
      enabled: environment.production,
      registrationStrategy: 'registerWhenStable:30000',
    }),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule { }
