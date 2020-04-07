import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {DiplomLayoutComponent} from './diplom/diplom-layout/diplom-layout.component';
import {PmComponent} from './diplom/diplom-pm/pm.component';
import {DiplomReviewerComponent} from './diplom/diplom-reviewer/diplom-reviewer.component';
import {DiplomNormocontrollerComponent} from './diplom/diplom-normocontroller/diplom-normocontroller.component';
import {DiplomChairmanComponent} from './diplom/diplom-chairman/diplom-chairman.component';
import {DiplomCommissionComponent} from './diplom/diplom-commission/diplom-commission.component';
import {DiplomOrderComponent} from './diplom/diplom-order/diplom-order.component';


const routes: Routes = [{
  path: '',
  component: DiplomLayoutComponent,
  children: [
    {
      path: 'pm',
      component: PmComponent,
    },
    {
      path: 'chairman',
      component: DiplomChairmanComponent,
    },
    {
      path: 'commission',
      component: DiplomCommissionComponent,
    },
    {
      path: 'reviewer',
      component: DiplomReviewerComponent,
    },
    {
      path: 'normcontroller',
      component: DiplomNormocontrollerComponent,
    },
    {
      path: 'diplomOrder',
      component: DiplomOrderComponent,
    },
  ]
}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
