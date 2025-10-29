import { Pipe, PipeTransform } from '@angular/core';
import { ConnectorStandard } from '../../../types';

@Pipe({
  name: 'connectorStandard',
  standalone: true
})
export class ConnectorStandardPipe implements PipeTransform {
  transform(value: ConnectorStandard): string {
    const labels: Record<ConnectorStandard, string> = {
      'AC_1P': 'AC 1-Phase',
      'AC_3P': 'AC 3-Phase',
      'DC': 'DC'
    };
    
    return labels[value] || value;
  }
}

