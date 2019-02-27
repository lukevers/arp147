class AsciiTable
    new: =>
        @headers = {}
        @rows = {}

    addHeader: (header) =>
        table.insert @headers, header

    addRow: (cols) =>
        table.insert @rows, cols

    render: =>
        colSize = {}

        for index, header in pairs @headers
            colSize[index] = #header

        for ri, cols in pairs @rows
            for index, col in pairs cols
                colSize[index] = #col if colSize[index] < #col

        rHeader = "| "
        rHeader2 = "+"
        for index, header in pairs @headers
            spaces = ""
            spaces ..= " " for i = 0, colSize[index] - #header

            rHeader ..= "#{header}#{spaces}| "
            rHeader2 ..= "-" for i = 0, colSize[index] + 1
            rHeader2 ..= "+"

        print rHeader2
        print rHeader
        print rHeader2

        for ri, cols in pairs @rows
            rRow = "| "
            for index, col in pairs cols
                spaces = ""
                spaces ..= " " for i = 0, colSize[index] - #col
                rRow ..= "#{col}#{spaces}| "
            print rRow

        print rHeader2
return AsciiTable!