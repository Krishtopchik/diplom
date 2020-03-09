import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {DiplomLayoutComponent} from './diplom/diplom-layout/diplom-layout.component';
import {PmComponent} from './diplom/diplom-pm/pm.component';


const routes: Routes = [{
  path: '',
  component: DiplomLayoutComponent,
  children: [
    {
      path: 'pm',
      component: PmComponent,
    }
  ]
}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
