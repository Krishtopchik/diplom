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

  changeTab($event) {
    this.tab = $event.target.getAttribute('rel');
    this.router.navigate([this.tab]);
  }
}