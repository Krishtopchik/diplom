import {Component, Input, OnInit, DoCheck} from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from '@angular/forms';
import {ValidationService} from '../../common/services/validation.service';
import {ToastrService} from 'ngx-toastr';
import {DiplomService} from '../../common/services/diplom.service';
import {DiplomModel} from '../../common/models/diplom.model';

@Component({
  selector: 'app-diplom-detail',
  templateUrl: './diplom-detail.component.html',
  styleUrls: ['./diplom-detail.component.scss']
})
export class DiplomDetailComponent implements OnInit, DoCheck {

  constructor(
    private fb: FormBuilder,
    private validationService: ValidationService,
    private diplomService: DiplomService,
    private toastr: ToastrService
  ) {
  }

  @Input() selectDiplomId: number;
  @Input() isDiplomSelect: boolean;
  @Input() infoForCreate: any;

  diplomSelect = false;
  diplomId = 0;
  tab = 'add';
  diplomForm: FormGroup;
  buttonTitleAdd = true;

  ngOnInit(): void {
    this.formInit();
  }

  ngDoCheck() {
    this.diplomSelect = this.isDiplomSelect;
    if (this.diplomSelect && this.selectDiplomId !== this.diplomId) {
      this.diplomId = this.selectDiplomId;
      this.tab = 'add';
      this.getDiplom(this.selectDiplomId);
    }
  }

  onSubmit() {
    this.validationService.validateAndSave(this.diplomForm, () => {
      const diplom: DiplomModel = this.diplomForm.getRawValue();
      diplom.Completion = +diplom.Completion;
      diplom.Score = +diplom.Score;
      diplom.Queuenumber = +diplom.Queuenumber;
      diplom.PmId = +diplom.PmId;
      diplom.NormcontrollerId = +diplom.NormcontrollerId;
      diplom.ReviewerId = +diplom.ReviewerId;
      diplom.ChairmanId = +diplom.ChairmanId;
      diplom.DiplomorderId = +diplom.DiplomorderId;
      diplom.SpecialtyId = +diplom.SpecialtyId;
      diplom.CommissionId = +diplom.CommissionId;
      if (this.buttonTitleAdd) {
        this.diplomService.createDiplom(diplom).subscribe(res => {
          this.toastr.success('Добавлен');
        });
      } else {
        this.diplomService.updateDiplom(diplom).subscribe(res => {
          this.toastr.success('Изменен');
        });
      }
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
    if (this.tab === 'add') {
      this.formInit();
      this.buttonTitleAdd = true;
    }
  }

  private getDiplom(id: number) {
    this.diplomService.getDiplomById(id).subscribe(res => {
      this.fillForm(res);
      this.buttonTitleAdd = false;
    });
  }

  private formInit() {
    this.diplomForm = this.fb.group({
      Id: [0],
      Fio: ['', Validators.required],
      Topic: ['', Validators.required],
      Completion: ['', Validators.required],
      Score: ['', Validators.required],
      Deadline: ['', Validators.required],
      Queuenumber: ['', Validators.required],
      PmId: [1, Validators.required],
      NormcontrollerId: [1, Validators.required],
      ReviewerId: [1, Validators.required],
      ChairmanId: [1, Validators.required],
      DiplomorderId: [1, Validators.required],
      SpecialtyId: [1, Validators.required],
      CommissionId: [1, Validators.required],
    });
  }

  private fillForm(item: DiplomModel) {
    this.diplomForm.patchValue({
      ...item
    });
  }
}
