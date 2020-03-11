import {Component, Input, OnInit} from '@angular/core';
import {DiplomService} from '../../common/services/diplom.service';
import {DiplomModel} from '../../common/models/diplom.model';
import {DiplomDataService} from '../../common/services/diplom-data.service';
import {DiplomInfoModel} from '../../common/models/diplomInfo.model';

@Component({
  selector: 'app-diplom-list',
  templateUrl: './diplom-list.component.html',
  styleUrls: ['./diplom-list.component.scss']
})
export class DiplomListComponent implements OnInit {

  @Input() diplomsList: DiplomModel[];
  @Input() infoAboutDiplom: DiplomInfoModel;

  constructor(
    private diplomDataService: DiplomDataService,
    private diplomService: DiplomService,
  ) {
  }

  ngOnInit(): void {
  }

  changeDiplom(id: number) {
    this.diplomDataService.selectDiplomId = id;
    this.diplomDataService.isDiplomSelect = true;
  }

  createTitleString(diplom: DiplomModel) {
    const diplomOrder = this.infoAboutDiplom.diplomorderList.find( (el) => el.Id === diplom.DiplomorderId);
    const specialty = this.infoAboutDiplom.specialtyList.find( (el) => el.Id === diplom.SpecialtyId);
    const pm = this.infoAboutDiplom.pmList.find( (el) => el.Id === diplom.PmId);
    return `[${diplomOrder.Name}, ${diplomOrder.Dateorder}]-${diplom.Queuenumber}-${specialty.Name}-${diplom.Fio}/${pm.Fio}`;
  }

  getNormcontroller(normcontrollerId: number) {
    const normcontroller = this.infoAboutDiplom.normcontrollerList.find( (el) => el.Id === normcontrollerId);
    return normcontroller.Fio;
  }

  getReviewer(reviewerId: number) {
    const reviewer = this.infoAboutDiplom.reviewerList.find( (el) => el.Id === reviewerId);
    return reviewer.Fio;
  }

  getCommission(commissionId: number) {
    const commission = this.infoAboutDiplom.commissionList.find( (el) => el.Id === commissionId);
    return commission.Fio;
  }

  getChairman(chairmanId: number) {
    const chairman = this.infoAboutDiplom.chairmanList.find( (el) => el.Id === chairmanId);
    return chairman.Fio;
  }

  deleteDiplom(id: number) {
    this.diplomService.deleteDiplom(id).subscribe(res => {
      this.diplomDataService.isDiplomsUpdate = true;
    });
  }
}
