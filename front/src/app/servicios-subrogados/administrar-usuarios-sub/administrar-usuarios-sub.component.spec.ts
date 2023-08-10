import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdministrarUsuariosSubComponent } from './administrar-usuarios-sub.component';

describe('AdministrarUsuariosSubComponent', () => {
  let component: AdministrarUsuariosSubComponent;
  let fixture: ComponentFixture<AdministrarUsuariosSubComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AdministrarUsuariosSubComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AdministrarUsuariosSubComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
