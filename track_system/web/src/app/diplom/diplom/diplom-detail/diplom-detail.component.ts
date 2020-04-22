import {Component, Input, OnInit, DoCheck, Output, EventEmitter} from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from '@angular/forms';
import {ValidationService} from '../../../common/services/validation.service';
import {ToastrService} from 'ngx-toastr';
import {DiplomService} from '../../../common/services/diplom.service';
import {DiplomModel} from '../../../common/models/diplom.model';
import {DiplomDataService} from '../../../common/services/diplom-data.service';
import {DiplomInfoModel} from '../../../common/models/diplomInfo.model';

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
    private toastr: ToastrService,
    private diplomDataService: DiplomDataService
  ) {
  }

  @Input() infoAboutDiplom: DiplomInfoModel;
  @Input() selectDiplomId: number;
  @Input() isDiplomSelect: boolean;

  diplomSelect = false;
  diplomId = 0;
  tab = 'add';
  diplomForm: FormGroup;
  diplomListFilterForm: FormGroup;
  buttonTitleAdd = true;

  ngOnInit(): void {
    this.formFilterInit();
    this.formInit();
  }

  ngDoCheck() {
    this.diplomSelect = this.isDiplomSelect;
    if (this.diplomSelect && this.selectDiplomId !== this.diplomId) {
      this.diplomId = this.selectDiplomId;
      this.tab = 'change';
      this.getDiplom(this.selectDiplomId);
      this.diplomSelect = false;
      this.diplomDataService.isDiplomSelect = false;
      this.diplomDataService.selectDiplomId = null;
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
      if (diplom.Deadline === '') {
        diplom.Deadline = null;
      }
      if (diplom.Execution === '') {
        diplom.Execution = null;
      }
      if (this.buttonTitleAdd) {
        this.diplomService.createDiplom(diplom).subscribe(res => {
          this.toastr.success('Добавлен');
        });
      } else {
        this.diplomService.updateDiplom(diplom).subscribe(res => {
          this.toastr.success('Изменен');
        });
      }
      this.diplomDataService.isDiplomsUpdate = true;
    }, {
      Fio: {
        required: 'Поле Студент не заполнено'
      },
      Topic: {
        required: 'Поле Тема не заполнено'
      },
      Type: {
        required: 'Не вырбран тип'
      },
    });
  }

  onFilterSubmit() {
    this.diplomDataService.diplomsFilter = true;
    const rows = this.diplomListFilterForm.getRawValue();
    const filter = {};
    Object.keys(rows).forEach(key => {
      if (rows[`${key}Check`]) {
        if (rows[key] === 'null') {
          rows[key] = 0;
        }
        filter[key] = rows[key] ;
      }
    });
    localStorage.setItem('filter', JSON.stringify(filter));
  }

  changeTab($event) {
    this.diplomId = null;
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
      Completion: [''],
      Score: [''],
      Deadline: [''],
      Queuenumber: [''],
      PmId: [0],
      NormcontrollerId: [0],
      ReviewerId: [0],
      ChairmanId: [0],
      DiplomorderId: [1],
      SpecialtyId: [1],
      CommissionId: [0],
      Execution: [''],
      Type: [0, Validators.required],
      CommissionComment: [''],
    });
  }

  formFilterInit() {
    this.diplomListFilterForm = this.fb.group({
      Fio: [''],
      Topic: [''],
      Completion: [''],
      Score: [''],
      Deadline: [''],
      Queuenumber: [''],
      PmId: [0],
      NormcontrollerId: [0],
      ReviewerId: [0],
      ChairmanId: [0],
      DiplomorderId: [1],
      SpecialtyId: [0],
      CommissionId: [0],
      Execution: [''],
      Type: [0],
      FioCheck: [false],
      TopicCheck: [false],
      CompletionCheck: [false],
      ScoreCheck: [false],
      DeadlineCheck: [false],
      QueuenumberCheck: [false],
      PmIdCheck: [false],
      NormcontrollerIdCheck: [false],
      ReviewerIdCheck: [false],
      ChairmanIdCheck: [false],
      DiplomorderIdCheck: [false],
      SpecialtyIdCheck: [false],
      CommissionIdCheck: [false],
      ExecutionCheck: [false],
      TypeCheck: [false]
    });
    const filterItem = JSON.parse(localStorage.getItem('filter'));
    if (filterItem) {
      Object.keys(filterItem).forEach(key => {
        const keyObj = `${key}Check`;
        filterItem[keyObj] = true;
      });
      this.diplomListFilterForm.patchValue({
        ...filterItem
      });
    }
  }

  private fillForm(item: DiplomModel) {
    this.diplomForm.patchValue({
      ...item
    });
  }

  removeFilter() {
    localStorage.removeItem('filter');
    this.formFilterInit();
    this.diplomDataService.isDiplomsUpdate = true;
  }

  clearForm() {
    this.diplomId = null;
    this.formInit();
    this.tab = 'add';
  }
}
