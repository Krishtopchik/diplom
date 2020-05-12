import {Component, DoCheck, OnInit} from '@angular/core';
import {Router} from '@angular/router';

@Component({
  selector: 'app-diplom-layout',
  templateUrl: './diplom-layout.component.html',
  styleUrls: ['./diplom-layout.component.scss']
})
export class DiplomLayoutComponent {

  tab = '';

  constructor(
    private router: Router,
  ) {

  }

  changeTab(rel) {
    this.tab = rel;
    this.router.navigate([rel]);
  }
}
