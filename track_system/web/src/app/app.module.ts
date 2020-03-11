import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DiplomComponent } from './diplom/diplom/diplom.component';
import { DiplomListComponent } from './diplom/diplom-list/diplom-list.component';
import { DiplomDetailComponent } from './diplom/diplom-detail/diplom-detail.component';
import { PmComponent } from './diplom/diplom-pm/pm.component';
import { DiplomLayoutComponent } from './diplom/diplom-layout/diplom-layout.component';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {HttpClientModule} from '@angular/common/http';
import {ValidationService} from './common/services/validation.service';
import {ToastContainerModule, ToastrModule} from 'ngx-toastr';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {AccordionConfig, AccordionModule} from 'ngx-bootstrap/accordion';

@NgModule({
  declarations: [
    AppComponent,
    DiplomComponent,
    DiplomListComponent,
    DiplomDetailComponent,
    PmComponent,
    DiplomLayoutComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot({positionClass: 'inline'}),
    ToastContainerModule,
    BrowserModule,
    AppRoutingModule,
    NgbModule,
    ReactiveFormsModule,
    HttpClientModule,
    FormsModule,
    AccordionModule,
  ],
  providers: [ValidationService, AccordionConfig],
  bootstrap: [AppComponent]
})
export class AppModule { }
