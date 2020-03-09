import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {DiplomService} from '../../common/services/diplom.service';
import {DiplomModel} from '../../common/models/diplom.model';
import {TeacherModel} from '../../common/models/teacher.model';
import {SpecialytyModel} from '../../common/models/specialyty.model';
import {DiplomorderModel} from '../../common/models/diplomorder.model';
import {zip} from 'rxjs';
import {validate} from 'codelyzer/walkerFactory/walkerFn';

@Component({
  selector: 'app-diplom',
  templateUrl: './diplom.component.html',
  styleUrls: ['./diplom.component.scss']
})
export class DiplomComponent implements OnInit {

  diplomForm = new FormGroup({
    Fio: new FormControl('', Validators.required),
    Topic: new FormControl('', Validators.required),
    Completion: new FormControl('', Validators.required),
    Score: new FormControl('', Validators.required),
    Deadline: new FormControl('', Validators.required),
    Queuenumber: new FormControl('', Validators.required),
    PmId: new FormControl('', Validators.required),
    NormcontrollerId: new FormControl('', Validators.required),
    ReviewerId: new FormControl('', Validators.required),
    ChairmanId: new FormControl('', Validators.required),
    DiplomorderId: new FormControl('', Validators.required),
    SpecialtyId: new FormControl('', Validators.required),
    CommissionId: new FormControl('', Validators.required),
  });

  diplomsList: DiplomModel[];
  infoForCreate: any;

  constructor(
    private diplomService: DiplomService,
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
