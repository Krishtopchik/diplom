import { Component, OnInit } from '@angular/core';
import {TeacherModel} from '../../common/models/teacher.model';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {DiplomService} from '../../common/services/diplom.service';
import {ToastrService} from 'ngx-toastr';
import {DiplomorderModel} from '../../common/models/diplomorder.model';

@Component({
  selector: 'app-diplom-order',
  templateUrl: './diplom-order.component.html',
  styleUrls: ['./diplom-order.component.scss']
})
export class DiplomOrderComponent implements OnInit {

  diplomOrderList: DiplomorderModel[];
  diplomOrderForm: FormGroup;
  buttonTitleAdd = true;
  tab = 'add';
  constructor(
    private diplomService: DiplomService,
    private fb: FormBuilder,
    private toastr: ToastrService,
  ) {
  }

  ngOnInit(): void {
    this.getReviewerList();
    this.formInit();
  }

  private getReviewerList() {
    this.diplomService.getDiplomorders().subscribe(res => {
      this.diplomOrderList = res;
    });
  }

  private formInit() {
    this.diplomOrderForm = this.fb.group({
      Id: [0],
      Name: ['', Validators.required],
      Dateorder: ['', Validators.required],
    });
  }

  onSubmit() {
    const diplomOrder: TeacherModel = this.diplomOrderForm.getRawValue();
    if (this.buttonTitleAdd) {
      this.diplomService.createDiplomorder(diplomOrder).subscribe(res => {
        this.toastr.success('Добавлен');
        this.getReviewerList();
      });
    } else {
      this.diplomService.updateeDiplomorder(diplomOrder).subscribe(res => {
        this.toastr.success('Обновлен');
        this.getReviewerList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }

  deleteDiplomOrder(id: number) {
    if (confirm('Удалить?')) {
      this.diplomService.deleteDiplomorder(id).subscribe(res => {
        this.toastr.success('Удален');
        this.getReviewerList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }

  changeDiplomOrder(id: number) {
    const item = this.diplomOrderList.find(el => el.Id === id);
    this.diplomOrderForm.patchValue({
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
