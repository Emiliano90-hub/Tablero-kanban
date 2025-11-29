import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateCardLayer } from './create-card-layer';

describe('CreateCardLayer', () => {
  let component: CreateCardLayer;
  let fixture: ComponentFixture<CreateCardLayer>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateCardLayer]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateCardLayer);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
