import {Component, OnInit} from '@angular/core';
import {DiplomService} from '../../common/services/diplom.service';
import {TeacherModel} from '../../common/models/teacher.model';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {ToastrService} from 'ngx-toastr';
import {DiplomModel} from '../../common/models/diplom.model';

@Component({
  selector: 'app-pm',
  templateUrl: './pm.component.html',
  styleUrls: ['./pm.component.scss']
})
export class PmComponent implements OnInit {

  pmList: TeacherModel[];
  pmForm: FormGroup;
  buttonTitleAdd = true;

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
    this.diplomService.getPms().subscribe(res => {
      this.pmList = res;
    });
  }

  private formInit() {
    this.pmForm = this.fb.group({
      Id: [0],
      Fio: ['', Validators.required],
    });
  }

  onSubmit() {
    const pm: TeacherModel = this.pmForm.getRawValue();
    if (this.buttonTitleAdd) {
      this.diplomService.createPm(pm).subscribe(res => {
        this.toastr.success('Добавлен');
        this.getPmList();
      });
    } else {
      this.diplomService.updateePm(pm).subscribe(res => {
        this.toastr.success('Обновлен');
        this.getPmList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
  }

  deletePm(id: number) {
    if (confirm('Удалить?')) {
      this.diplomService.deletePm(id).subscribe(res => {
        this.toastr.success('Удален');
        this.getPmList();
      });
    }
    this.formInit();
    this.buttonTitleAdd = true;
  }

  changePm(id: number) {
    const item = this.pmList.find(el => el.Id === id);
    this.pmForm.patchValue({
      ...item
    });
    this.buttonTitleAdd = false;
  }

  changeTab(ev) {
    this.formInit();
    this.buttonTitleAdd = true;
  }
}
