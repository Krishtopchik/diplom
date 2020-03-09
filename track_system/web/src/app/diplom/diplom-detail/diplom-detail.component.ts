import {Component, Input, OnInit} from '@angular/core';
import {FormGroup} from '@angular/forms';
import {SuperForm} from 'angular-super-validator';
import {ValidationService} from '../../common/services/validation.service';
import {ToastrService} from 'ngx-toastr';
import {DiplomService} from '../../common/services/diplom.service';
import {DiplomModel} from '../../common/models/diplom.model';

@Component({
  selector: 'app-diplom-detail',
  templateUrl: './diplom-detail.component.html',
  styleUrls: ['./diplom-detail.component.scss']
})
export class DiplomDetailComponent implements OnInit {

  constructor(
    private validationService: ValidationService,
    private diplomService: DiplomService,
    private toastr: ToastrService
  ) {
  }

  @Input() diplomForm: FormGroup;
  @Input() infoForCreate: any;

  tab = 'add';

  ngOnInit(): void {
  }

  onSubmit() {
    this.validationService.validateAndSave(this.diplomForm, () => {
      const diplom: DiplomModel = this.diplomForm.getRawValue();
      console.log(diplom)
      diplom.Completion = +diplom.Completion;
      diplom.Score = + diplom.Score;
      diplom.Queuenumber = +diplom.Queuenumber;
      diplom.PmId = +diplom.PmId;
      diplom.NormcontrollerId = +diplom.NormcontrollerId;
      diplom.ReviewerId = +diplom.ReviewerId;
      diplom.ChairmanId = +diplom.ChairmanId;
      diplom.DiplomorderId = +diplom.DiplomorderId;
      diplom.SpecialtyId = +diplom.SpecialtyId;
      diplom.CommissionId = +diplom.CommissionId;
      this.diplomService.createDiplom(diplom).subscribe(res => {
        this.toastr.success('Добавлен');
      });
    }, {
      Fio: {
        required: 'Поле не заполнено'
      },
      Topic: {
        required: 'Поле не заполнено'
      },
      Completion: {
        required: 'Поле не заполнено'
      },
      Score: {
        required: 'Поле не заполнено'
      },
      Deadline: {
        required: 'Поле не заполнено'
      },
      Queuenumber: {
        required: 'Поле не заполнено'
      },
      PmId: {
        required: 'Поле не заполнено'
      },
      NormcontrollerId: {
        required: 'Поле не заполнено'
      },
      ReviewerId: {
        required: 'Поле не заполнено'
      },
      ChairmanId: {
        required: 'Поле не заполнено'
      },
      DiplomorderId: {
        required: 'Поле не заполнено'
      },
      SpecialtyId: {
        required: 'Поле не заполнено'
      },
      CommissionId: {
        required: 'Поле не заполнено'
      }
    });
  }

  changeTab($event) {
    this.tab = $event.target.getAttribute('rel');
    console.log(this.tab)
  }
}
