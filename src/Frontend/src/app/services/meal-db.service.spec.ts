import { TestBed } from '@angular/core/testing';

import { MealDbService } from './meal-db.service';

describe('MealDbService', () => {
  let service: MealDbService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MealDbService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
