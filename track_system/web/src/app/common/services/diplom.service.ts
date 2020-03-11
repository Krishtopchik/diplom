import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {DiplomModel} from '../models/diplom.model';
import {TeacherModel} from '../models/teacher.model';
import {DiplomorderModel} from '../models/diplomorder.model';
import {SpecialytyModel} from '../models/specialyty.model';

@Injectable({
  providedIn: 'root'
})
export class DiplomService {
  url = 'http://localhost:8185/api/';

  constructor(private http: HttpClient) {
  }

  getAllDiploms(): Observable<DiplomModel[]> {
    return this.http.get<DiplomModel[]>(
      `${this.url}/diploms`
    );
  }

  getDiplomById(id): Observable<DiplomModel> {
    return this.http.get<DiplomModel>(
      `${this.url}/diploms/${id}`
    );
  }

  updateDiplom(diplom: DiplomModel): Observable<any> {
    return this.http.put<DiplomModel>(`${this.url}diploms`, JSON.stringify(diplom));
  }

  deleteDiplom(id: number): Observable<any> {
    return this.http.delete<any>(
      `${this.url}/diploms/${id}`
    );
  }

  createDiplom(diplom: DiplomModel): Observable<DiplomModel> {
    return this.http.post<DiplomModel>(`${this.url}diploms`, JSON.stringify(diplom));
  }

  getChairmans(): Observable<TeacherModel[]> {
    return this.http.get<TeacherModel[]>(
      `${this.url}/chairmans`
    );
  }

  getCommissions(): Observable<TeacherModel[]> {
    return this.http.get<TeacherModel[]>(
      `${this.url}/commissions`
    );
  }

  getDiplomorders(): Observable<DiplomorderModel[]> {
    return this.http.get<DiplomorderModel[]>(
      `${this.url}/diplomorders`
    );
  }

  getNormcontrollers(): Observable<TeacherModel[]> {
    return this.http.get<TeacherModel[]>(
      `${this.url}/normcontrollers`
    );
  }

  getPms(): Observable<TeacherModel[]> {
    return this.http.get<TeacherModel[]>(
      `${this.url}/pms`
    );
  }

  getReviewers(): Observable<TeacherModel[]> {
    return this.http.get<TeacherModel[]>(
      `${this.url}/reviewers`
    );
  }

  getSpecialtys(): Observable<SpecialytyModel[]> {
    return this.http.get<SpecialytyModel[]>(
      `${this.url}/specialtys`
    );
  }
}
