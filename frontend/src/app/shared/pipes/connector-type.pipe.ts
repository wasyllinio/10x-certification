import { Pipe, PipeTransform } from '@angular/core';
import { ConnectorType } from '../../../types';

@Pipe({
  name: 'connectorType',
  standalone: true
})
export class ConnectorTypePipe implements PipeTransform {
  transform(value: ConnectorType): string {
    const labels: Record<ConnectorType, string> = {
      'CCS': 'CCS',
      'Type2': 'Type 2',
      'Chademo': 'CHAdeMO'
    };
    
    return labels[value] || value;
  }
}

