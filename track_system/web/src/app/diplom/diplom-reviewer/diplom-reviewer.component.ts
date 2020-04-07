import { Component, OnInit } from '@angular/core';
import {TeacherModel} from '../../common/models/teacher.model';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {DiplomService} from '../../common/services/diplom.service';
import {ToastrService} from 'ngx-toastr';

@Component({
  selector: 'app-diplom-reviewer',
  templateUrl: './diplom-reviewer.component.html',
  styleUrls: ['./diplom-reviewer.component.scss']
})
export class DiplomReviewerComponent implements OnInit {

  reviewerList: TeacherModel[];
  reviewerForm: FormGroup;
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
    this.diplomService.getReviewers().subscribe(res => {
      this.reviewerList = res;
    });
  }

  private formInit() {
    this.reviewerForm = this.fb.group({
      Id: [0],
      Fio: ['', Validators.required],
    });
  }

  onSubmit() {
    const reviewer: TeacherModel = this.reviewerForm.getRawValue();
    if (this.buttonTitleAdd) {
      this.diplomService.createReviewer(reviewer).subscribe(res => {
        this.toastr.success('Добавлен');
        this.getReviewerList();
      });
    } else {
      this.diplomService.updateeReviewer(reviewer).subscribe(res => {
        this.toastr.success('Обновлен');
        this.getReviewerList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }

  deleteReviewer(id: number) {
    if (confirm('Удалить?')) {
      this.diplomService.deleteReviewer(id).subscribe(res => {
        this.toastr.success('Удален');
        this.getReviewerList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
    this.tab = 'add';
  }

  changeReviewer(id: number) {
    const item = this.reviewerList.find(el => el.Id === id);
    this.reviewerForm.patchValue({
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
