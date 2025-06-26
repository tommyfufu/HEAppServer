import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../api-service.service';
import { Record } from '../models/patient.model';
import { NgIf, NgFor } from '@angular/common';

@Component({
  selector: 'app-game-record',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './game-record.component.html',
  styleUrl: './game-record.component.css'
})

export class GameRecordComponent {
  @Input() userId?: string;     // 從父元件傳入
  records: Record[] = [];
  loading = false;
  error: string | null = null;

  constructor(private api: ApiService) {}

  ngOnChanges(changes: SimpleChanges) {
    if (changes['userId'] && this.userId) {
      this.fetchRecords(this.userId);
    }
  }

  private fetchRecords(userId: string) {
    this.loading = true;
    this.error = null;
    this.api.getRecords(userId).subscribe({
      next: (recs) => {
        this.records = recs;
        this.loading = false;
      },
      error: (err) => {
        console.error('Failed to load records', err);
        this.error = '無法讀取遊戲紀錄';
        this.loading = false;
      }
    });
  }
}
