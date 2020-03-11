import {Injectable} from '@angular/core';

@Injectable({ providedIn: 'root' })
export class DiplomDataService {
  selectDiplomId: number;
  isDiplomSelect = false;
  isDiplomsUpdate = false;
  diplomsFilter = false;
}
