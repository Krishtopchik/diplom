import { Component, OnInit } from '@angular/core';
import {DiplomService} from '../../common/services/diplom.service';
import {DiplomModel} from '../../common/models/diplom.model';

@Component({
  selector: 'app-diplom-list',
  templateUrl: './diplom-list.component.html',
  styleUrls: ['./diplom-list.component.scss']
})
export class DiplomListComponent implements OnInit {

  constructor(
    private diplomService: DiplomService
    ) { }
  diplomList: DiplomModel[];
  ngOnInit(): void {
    this.getAllDiploms();
  }

  getAllDiploms() {
    this.diplomService.getAllDiploms().subscribe(res => {
      this.diplomList = res;
    })
  }
}
