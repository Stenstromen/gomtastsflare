package resource

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func ResourceReport(cmd *cobra.Command, args []string) error {
	filename, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var report Report

	gz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Print(err)
	}
	defer gz.Close()

	err = json.NewDecoder(gz).Decode(&report)
	if err != nil {
		fmt.Print(err)
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("  ")
	table.SetNoWhiteSpace(true)
	table.SetColumnSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Organization", "Start Datetime", "End Datetime", "Contact Info", "Report ID", "Type", "Policy", "Domain", "MX Host", "Successes", "Failures"})

	for _, policy := range report.Policies {
		for i, policyString := range policy.Policy.PolicyString {
			row := []string{
				report.OrganizationName,
				report.DateRange.StartDatetime,
				report.DateRange.EndDatetime,
				report.ContactInfo,
				report.ReportID,
				policy.Policy.PolicyType,
				policyString,
				policy.Policy.PolicyDomain,
				strings.Join(policy.Policy.MXHost, ", "),
				fmt.Sprintf("%d", policy.Summary.TotalSuccessfulSessionCount),
				fmt.Sprintf("%d", policy.Summary.TotalFailureSessionCount),
			}
			if i == -1 {
				row[5] = policy.Policy.PolicyType
				row[7] = policy.Policy.PolicyDomain
				row[8] = strings.Join(policy.Policy.MXHost, ", ")
			}
			table.Append(row)
		}
	}

	table.Render()

	return nil
}
