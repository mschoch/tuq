package parser

func concatStringSlices(old1, old2 []string) []string {
    newslice := make([]string, len(old1)+len(old2))
    copy(newslice, old1)
    copy(newslice[len(old1):], old2)
    return newslice
}