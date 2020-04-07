import { Component, OnInit } from '@angular/core';
import {TeacherModel} from '../../common/models/teacher.model';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {DiplomService} from '../../common/services/diplom.service';
import {ToastrService} from 'ngx-toastr';

@Component({
  selector: 'app-diplom-commission',
  templateUrl: './diplom-commission.component.html',
  styleUrls: ['./diplom-commission.component.scss']
})
export class DiplomCommissionComponent implements OnInit {

  commissionList: TeacherModel[];
  commissionForm: FormGroup;
  buttonTitleAdd = true;
  tab = 'add';
  constructor(
    private diplomService: DiplomService,
    private fb: FormBuilder,
    private toastr: ToastrService,
  ) {
  }

  ngOnInit(): void {
    this.getPmList();
    this.formInit();
  }

  private getPmList() {
    this.diplomService.getCommissions().subscribe(res => {
      this.commissionList = res;
    });
  }

  private formInit() {
    this.commissionForm = this.fb.group({
      Id: [0],
      Fio: ['', Validators.required],
    });
  }

  onSubmit() {
    const commission: TeacherModel = this.commissionForm.getRawValue();
    if (this.buttonTitleAdd) {
      this.diplomService.createCommission(commission).subscribe(res => {
        this.toastr.success('Добавлен');
        this.getPmList();
      });
    } else {
      this.diplomService.updateeCommission(commission).subscribe(res => {
        this.toastr.success('Обновлен');
        this.getPmList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }

  deleteChairman(id: number) {
    if (confirm('Удалить?')) {
      this.diplomService.deleteCommission(id).subscribe(res => {
        this.toastr.success('Удален');
        this.getPmList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }

  changeChairman(id: number) {
    const item = this.commissionList.find(el => el.Id === id);
    this.commissionForm.patchValue({
      ...item
    });
    this.buttonTitleAdd = false;
    this.tab = 'change';
  }

  changeTab(ev) {
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }
}
