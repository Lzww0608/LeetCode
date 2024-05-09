type TextEditor struct {
    pre []byte
    suf []byte
}


func Constructor() TextEditor {
    return TextEditor{}
}


func (this *TextEditor) AddText(text string)  {
    for i := range text {
        this.pre = append(this.pre, text[i])
    }
}


func (this *TextEditor) DeleteText(k int) (ans int) {
    for i := 0; i < k && len(this.pre) > 0; i++ {
        this.pre = this.pre[:len(this.pre)-1]
        ans++
    }
    return 
}


func (this *TextEditor) CursorLeft(k int) string {
    for i := 0; i < k && len(this.pre) > 0; i++ {
        tmp := this.pre[len(this.pre)-1]
        this.pre = this.pre[:len(this.pre)-1]
        this.suf = append(this.suf, tmp)
    }
    return string(this.pre[max(0, len(this.pre)-10):len(this.pre)])
}


func (this *TextEditor) CursorRight(k int) string {
    for i := 0; i < k && len(this.suf) > 0; i++ {
        tmp := this.suf[len(this.suf)-1]
        this.suf = this.suf[:len(this.suf)-1]
        this.pre = append(this.pre, tmp)
    }
    return string(this.pre[max(0, len(this.pre)-10):len(this.pre)])
}


/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */