import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DiplomComponent } from './diplom/diplom/diplom.component';
import { DiplomListComponent } from './diplom/diplom/diplom-list/diplom-list.component';
import { DiplomDetailComponent } from './diplom/diplom/diplom-detail/diplom-detail.component';
import { PmComponent } from './diplom/diplom-pm/pm.component';
import { DiplomLayoutComponent } from './diplom/diplom-layout/diplom-layout.component';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {HttpClientModule} from '@angular/common/http';
import {ValidationService} from './common/services/validation.service';
import {ToastContainerModule, ToastrModule} from 'ngx-toastr';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {AccordionConfig, AccordionModule} from 'ngx-bootstrap/accordion';
import { DiplomReviewerComponent } from './diplom/diplom-reviewer/diplom-reviewer.component';
import { DiplomNormocontrollerComponent } from './diplom/diplom-normocontroller/diplom-normocontroller.component';
import { DiplomChairmanComponent } from './diplom/diplom-chairman/diplom-chairman.component';
import { DiplomCommissionComponent } from './diplom/diplom-commission/diplom-commission.component';
import { DiplomOrderComponent } from './diplom/diplom-order/diplom-order.component';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {MatSelectModule} from '@angular/material/select';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatRadioModule} from '@angular/material/radio';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatExpansionModule} from '@angular/material/expansion';
import { DiplomIdComponent } from './diplom/diplom/diplom-id/diplom-id.component';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatListModule} from '@angular/material/list';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatDatepickerModule} from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import {IConfig, NgxMaskModule} from 'ngx-mask';

// export const options: Partial<IConfig> | (() => Partial<IConfig>);

@NgModule({
  declarations: [
    AppComponent,
    DiplomComponent,
    DiplomListComponent,
    DiplomDetailComponent,
    PmComponent,
    DiplomLayoutComponent,
    DiplomReviewerComponent,
    DiplomNormocontrollerComponent,
    DiplomChairmanComponent,
    DiplomCommissionComponent,
    DiplomOrderComponent,
    DiplomIdComponent,
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
    MatButtonModule,
    MatInputModule,
    MatSelectModule,
    MatCheckboxModule,
    MatRadioModule,
    MatFormFieldModule,
    MatExpansionModule,
    MatGridListModule,
    MatListModule,
    MatProgressSpinnerModule,
    MatDatepickerModule,
    MatNativeDateModule,
    NgxMaskModule.forRoot(),
  ],
  providers: [ValidationService, AccordionConfig],
  bootstrap: [AppComponent]
})
export class AppModule { }
