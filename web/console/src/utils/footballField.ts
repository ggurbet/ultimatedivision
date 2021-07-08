export interface VisibilityStyles {
    visible: boolean;
    get style(): string;
}
export class TriangleStyle implements VisibilityStyles {
    constructor(
        public visible: boolean = false
    ) { }
    get style(): string {
        return this.visible ? 'rotate(-90deg)' : 'rotate(0deg)';
    }
}
export class ListStyle implements VisibilityStyles {
    constructor(
        public visible: boolean = false
    ) { }
    get style(): string {
        return this.visible ? '0' : '90px';
    }
}

export class FootballCardStyle implements VisibilityStyles {
    constructor (
        public visible: boolean
        ) { }
        get style(): string {
            return this.visible ? 'block' : 'none';
        }
}