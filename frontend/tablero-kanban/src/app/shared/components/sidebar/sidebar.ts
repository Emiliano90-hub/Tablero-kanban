import { Component, EventEmitter, Output } from '@angular/core';
import { SidebarOption } from '../enums/sidebar-option.enum';
@Component({
  standalone: true,
  selector: 'app-sidebar',
  imports: [],
  templateUrl: './sidebar.html',
  styleUrl: './sidebar.css',
})
export class Sidebar {
  @Output() selectComponent = new EventEmitter<SidebarOption>();

  option: SidebarOption = SidebarOption.None;
  SidebarOption = SidebarOption;

  select(component: SidebarOption) {
    this.option = this.option === component ? SidebarOption.None : component;
    this.selectComponent.emit(this.option);
  }

}
