export interface DropDownStyles {
    state: boolean;
    get style(): string;
}
export class TriangleStyle implements DropDownStyles {
    constructor(
        public state: boolean = false
    ) { }
    get style() {
        let style = this.state ? 'rotate(-90deg)' : 'rotate(0deg)';
        return style;
    }
}
export class ListStyle implements DropDownStyles {
    constructor(
        public state: boolean = false
    ) { }
    get style() {
        let style = this.state ? '0' : '90px';
        return style;
    }
}