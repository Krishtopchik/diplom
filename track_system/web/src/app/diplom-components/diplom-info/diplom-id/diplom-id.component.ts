import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {DiplomService} from '../../../common/services/diplom.service';
import {DiplomModel} from '../../../common/models/diplom.model';
import {zip} from 'rxjs';
import { timer } from 'rxjs';

@Component({
  selector: 'app-diplom-id',
  templateUrl: './diplom-id.component.html',
  styleUrls: ['./diplom-id.component.scss']
})
export class DiplomIdComponent implements OnInit {

  diplomId: number;
  pm = '';
  normcontroller = '';
  reviewer = '';
  chairman = '';
  diplomOrder = '';
  specialty = '';
  commission = '';
  diplom = new DiplomModel();
  subscribeTimer: any;
  timeLeft = 70;
  timerText: string;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private diplomService: DiplomService
  ) {
  }

  ngOnInit(): void {
    this.diplomId = this.route.snapshot.params.id;
    this.getDiplom();
  }

  getDiplom() {
    this.diplomService.getDiplomById(this.diplomId).subscribe(res => {
      if (res) {
        this.diplom = res;
        this.getDiplomInfo(res);
      }
    });
  }

  getDiplomInfo(diplom: DiplomModel) {
    zip(
      this.getPmList(),
      this.getNormcontrollerList(),
      this.getReviewerList(),
      this.getChairmanList(),
      this.getDiplomorderList(),
      this.getSpecialtyList(),
      this.getCommissionList(),
    ).subscribe(([pm, normcontroller, reviewer, chairman, diplomorder, specialty, commission]) => {
      this.pm = pm.find(el => el.Id === diplom.PmId).Fio;
      this.normcontroller = normcontroller.find(el => el.Id === diplom.NormcontrollerId).Fio;
      this.reviewer = reviewer.find(el => el.Id === diplom.ReviewerId).Fio;
      this.chairman = chairman.find(el => el.Id === diplom.ChairmanId).Fio;
      const diplomOrderModel = diplomorder.find(el => el.Id === diplom.DiplomorderId);
      this.diplomOrder = `${diplomOrderModel.Name} ${this.strToDate(diplomOrderModel.Dateorder)}`;
      this.specialty = specialty.find(el => el.Id === diplom.SpecialtyId).Name;
      this.commission = commission.find(el => el.Id === diplom.CommissionId).Fio;
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

  onMain() {
    this.router.navigate(['/']);
  }

  start() {
    const source = timer(1000, 1000);
    source.subscribe(val => {
      if (val < this.timeLeft) {
        this.subscribeTimer = this.timeLeft - val;
        this.timerText = this.secToMin(this.subscribeTimer);
      } else {
        this.timerText = 'конец';
      }
    });
  }

  secToMin(sec: string) {
    const min = Math.floor(+sec / 60);
    let seconds = (+sec - min * 60).toString();
    if (+seconds < 10) {
      seconds = `0${seconds}`;
    }
    return `${min}:${seconds}`;
  }

  strToDate(str: string) {
    const date = new Date(str);
    const day = date.getDate();
    const month = date.getMonth() + 1;
    const year = date.getFullYear();
    return `${day < 10 ? `0${day}` : day}.${month < 10 ? `0${month}` : month}.${year}`;
  }
}
