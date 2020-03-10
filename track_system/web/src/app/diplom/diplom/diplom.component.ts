import {Component, DoCheck, OnInit, ViewChild} from '@angular/core';
import {DiplomService} from '../../common/services/diplom.service';
import {DiplomModel} from '../../common/models/diplom.model';
import {zip} from 'rxjs';
import {DiplomDataService} from '../../common/services/diplom-data.service';

@Component({
  selector: 'app-diplom',
  templateUrl: './diplom.component.html',
  styleUrls: ['./diplom.component.scss']
})
export class DiplomComponent implements OnInit, DoCheck {

  isDiplomSelect = false;
  selectDiplomId: number;
  diplomsList: DiplomModel[];
  infoForCreate: any;
  constructor(
    private diplomService: DiplomService,
    private diplomDataService: DiplomDataService,
  ) {
  }

  ngOnInit() {
    this.getDiplomsList();
    zip(
      this.getPmList(),
      this.getNormcontrollerList(),
      this.getReviewerList(),
      this.getChairmanList(),
      this.getDiplomorderList(),
      this.getSpecialtyList(),
      this.getCommissionList(),
    ).subscribe(([pm, normcontroller, reviewer, chairman, diplomorder, specialty, commission]) => {
      this.infoForCreate = {
        pmList: pm,
        normcontrollerList: normcontroller,
        reviewerList: reviewer,
        chairmanList: chairman,
        diplomorderList: diplomorder,
        specialtyList: specialty,
        commissionList: commission,
      };
    });
  }

  ngDoCheck() {
    this.selectDiplomId = this.diplomDataService.selectDiplomId;
    this.isDiplomSelect = this.diplomDataService.isDiplomSelect;
  }

  private getDiplomsList() {
    this.diplomService.getAllDiploms().subscribe(res => {
      this.diplomsList = res;
    });
  }

  private getPmList() {
    return this.diplomService.getPms();
  }

  private getNormcontrollerList() {
    return this.diplomService.getNormcontrollers();
  }

  private getReviewerList() {
    return this.diplomService.getReviewers();
  }

  private getChairmanList() {
    return this.diplomService.getChairmans();
  }

  private getDiplomorderList() {
    return this.diplomService.getDiplomorders();
  }

  private getSpecialtyList() {
    return this.diplomService.getSpecialtys();
  }

  private getCommissionList() {
    return this.diplomService.getCommissions();
  }
}
