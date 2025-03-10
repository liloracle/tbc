// Returns if the two items are equal, or if both are null / undefined.
export function equalsOrBothNull<T>(a: T, b: T, comparator?: (_a: NonNullable<T>, _b: NonNullable<T>) => boolean): boolean {
  if (a == null && b == null)
    return true;

  if (a == null || b == null)
    return false;

  return (comparator || ((_a: NonNullable<T>, _b: NonNullable<T>) => a == b))(a!, b!);
}

export function sum(arr: Array<number>): number {
  return arr.reduce((total, cur) => total + cur, 0);
}

// Returns the index of maximum value, or null if empty.
export function maxIndex(arr: Array<number>): number | null {
	return arr.reduce((cur, v, i, arr) => v > arr[cur] ? i : cur, 0);
}

// Returns a new array containing only elements present in both a and b.
export function arrayEquals<T>(a: Array<T>, b: Array<T>, comparator?: (a: T, b: T) => boolean): boolean {
	comparator = comparator || ((a: T, b: T) => a == b);
	return a.length == b.length && a.every((val, i) => comparator!(val, b[i]));
}

// Returns a new array containing only elements present in both a and b.
export function intersection<T>(a: Array<T>, b: Array<T>): Array<T> {
  return a.filter(value => b.includes(value));
}

// Splits an array into buckets, where elements are placed in the same bucket if the
// toString function returns the same value.
export function bucket<T>(arr: Array<T>, toString: (val: T) => string): Record<string, Array<T>> {
	const buckets: Record<string, Array<T>> = {};
	arr.forEach(val => {
		const valString = toString(val);
		if (buckets[valString]) {
			buckets[valString].push(val);
		} else {
			buckets[valString] = [val];
		}
	});
	return buckets;
}

export function stDevToConf90(stDev: number, N: number) {
  return 1.645 * stDev / Math.sqrt(N);
}

export async function wait(ms: number): Promise<void> {
  return new Promise(resolve => setTimeout(resolve, ms));
}

// Only works for numeric enums
export function getEnumValues<E>(enumType: any): Array<E> {
  return Object.keys(enumType)
      .filter(key => !isNaN(Number(enumType[key])))
      .map(key => parseInt(enumType[key]) as unknown as E);
}

// Whether a click event was a right click.
export function isRightClick(event: MouseEvent): boolean {
  return event.button == 2;
}

// Converts from '#ffffff' --> 'rgba(255, 255, 255, alpha)'
export function hexToRgba(hex: string, alpha: number): string {
	if (/^#([A-Fa-f0-9]{3}){1,2}$/.test(hex)) {
		let parts = hex.substring(1).split('');
		if (parts.length == 3) {
			parts = [parts[0], parts[0], parts[1], parts[1], parts[2], parts[2]];
		}
		const c: any = '0x' + parts.join('');
		return 'rgba(' + [(c>>16)&255, (c>>8)&255, c&255].join(',') + ',' + alpha + ')';
	}
	throw new Error('Invalid hex color: ' + hex);
}

export function camelToSnakeCase(str: string): string {
	let result = str.replace(/[A-Z]/g, letter => `_${letter.toLowerCase()}`);
	if (result.startsWith('_')) {
		result = result.substring(1);
	}
	return result;
}

export function downloadJson(json: any, fileName: string) {
	const dataStr = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(json, null, 2));
	const downloadAnchorNode = document.createElement('a');
	downloadAnchorNode.setAttribute("href", dataStr);
	downloadAnchorNode.setAttribute("download", fileName);
	document.body.appendChild(downloadAnchorNode); // required for firefox
	downloadAnchorNode.click();
	downloadAnchorNode.remove();
}
