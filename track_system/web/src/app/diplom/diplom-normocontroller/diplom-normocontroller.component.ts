import { Component, OnInit } from '@angular/core';
import {TeacherModel} from '../../common/models/teacher.model';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {DiplomService} from '../../common/services/diplom.service';
import {ToastrService} from 'ngx-toastr';

@Component({
  selector: 'app-diplom-normocontroller',
  templateUrl: './diplom-normocontroller.component.html',
  styleUrls: ['./diplom-normocontroller.component.scss']
})
export class DiplomNormocontrollerComponent implements OnInit {

  normocontrollerList: TeacherModel[];
  normocontrollerForm: FormGroup;
  buttonTitleAdd = true;
  tab = 'add';
  constructor(
    private diplomService: DiplomService,
    private fb: FormBuilder,
    private toastr: ToastrService,
  ) {
  }

  ngOnInit(): void {
    this.getNormocontrollerList();
    this.formInit();
  }

  private getNormocontrollerList() {
    this.diplomService.getNormcontrollers().subscribe(res => {
      this.normocontrollerList = res;
    });
  }

  private formInit() {
    this.normocontrollerForm = this.fb.group({
      Id: [0],
      Fio: ['', Validators.required],
    });
  }

  onSubmit() {
    const normocontroller: TeacherModel = this.normocontrollerForm.getRawValue();
    if (this.buttonTitleAdd) {
      this.diplomService.createNormcontroller(normocontroller).subscribe(res => {
        this.toastr.success('Добавлен');
        this.getNormocontrollerList();
      });
    } else {
      this.diplomService.updateeNormcontroller(normocontroller).subscribe(res => {
        this.toastr.success('Обновлен');
        this.getNormocontrollerList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }

  deleteNormocontroller(id: number) {
    if (confirm('Удалить?')) {
      this.diplomService.deleteNormcontroller(id).subscribe(res => {
        this.toastr.success('Удален');
        this.getNormocontrollerList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }

  changeNormocontroller(id: number) {
    const item = this.normocontrollerList.find(el => el.Id === id);
    this.normocontrollerForm.patchValue({
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
