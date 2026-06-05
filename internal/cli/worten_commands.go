package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/emmassist-co/worten-pp-cli/internal/worten"
)

func newWortenResolveCmd(flags *rootFlags) *cobra.Command {
	var raw bool
	cmd := &cobra.Command{
		Use:   "resolve <product-url-or-id>",
		Short: "Resolve a Worten product URL or slug to a canonical product identifier",
		Args: func(cmd *cobra.Command, args []string) error {
			if flags.dryRun {
				return nil
			}
			if len(args) != 1 {
				return usageErr(fmt.Errorf("resolve requires a product URL or product UUID"))
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if dryRunOK(flags) {
				return flags.printJSON(cmd, map[string]any{"dry_run": true, "command": "resolve"})
			}
			svc, err := newWortenService(flags)
			if err != nil {
				return err
			}
			result, err := svc.Resolve(cmd.Context(), args[0], raw)
			if err != nil {
				return classifyAPIError(err, flags)
			}
			return flags.printJSON(cmd, result)
		},
	}
	cmd.Flags().BoolVar(&raw, "raw", false, "Return raw cache/update details")
	return cmd
}

func newWortenProductCmd(flags *rootFlags) *cobra.Command {
	var raw bool
	cmd := oneInputWortenCmd(flags, "product", "Fetch and normalize a Worten product", func(cmd *cobra.Command, input string) (any, error) {
		svc, err := newWortenService(flags)
		if err != nil {
			return nil, err
		}
		return svc.Product(cmd.Context(), input, raw)
	})
	cmd.Flags().BoolVar(&raw, "raw", false, "Return the raw product details payload")
	return cmd
}

func newWortenBuyerCmd(flags *rootFlags) *cobra.Command {
	var raw bool
	cmd := oneInputWortenCmd(flags, "buyer", "Fetch and normalize the buyer view for a Worten product", func(cmd *cobra.Command, input string) (any, error) {
		svc, err := newWortenService(flags)
		if err != nil {
			return nil, err
		}
		return svc.Buyer(cmd.Context(), input, raw)
	})
	cmd.Flags().BoolVar(&raw, "raw", false, "Return the raw details/specifications payloads")
	return cmd
}

func newWortenSpecsCmd(flags *rootFlags) *cobra.Command {
	var raw bool
	cmd := oneInputWortenCmd(flags, "specs", "Fetch Worten product technical specifications", func(cmd *cobra.Command, input string) (any, error) {
		svc, err := newWortenService(flags)
		if err != nil {
			return nil, err
		}
		return svc.Specs(cmd.Context(), input, raw)
	})
	cmd.Flags().BoolVar(&raw, "raw", false, "Return the raw technical specifications payload")
	return cmd
}

func newWortenStockCmd(flags *rootFlags) *cobra.Command {
	var raw bool
	var postalCode string
	var radius int
	cmd := oneInputWortenCmd(flags, "stock", "Fetch normalized Worten stock context for a product", func(cmd *cobra.Command, input string) (any, error) {
		svc, err := newWortenService(flags)
		if err != nil {
			return nil, err
		}
		return svc.Stock(cmd.Context(), input, worten.StockOptions{
			PostalCode: postalCode,
			RadiusKm:   radius,
		}, raw)
	})
	cmd.Flags().BoolVar(&raw, "raw", false, "Return the raw product/store-search payloads")
	cmd.Flags().StringVar(&postalCode, "postal-code", "", "Postal code for nearby-store lookup")
	cmd.Flags().IntVar(&radius, "radius", 20, "Nearby-store search radius in km")
	return cmd
}

func newWortenSuggestCmd(flags *rootFlags) *cobra.Command {
	var raw bool
	var max int
	cmd := &cobra.Command{
		Use:   "suggest <query>",
		Short: "Fetch Worten search suggestions",
		Args: func(cmd *cobra.Command, args []string) error {
			if flags.dryRun {
				return nil
			}
			if len(args) == 0 {
				return usageErr(fmt.Errorf("suggest requires a query"))
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if dryRunOK(flags) {
				return flags.printJSON(cmd, map[string]any{"dry_run": true, "command": "suggest"})
			}
			svc, err := newWortenService(flags)
			if err != nil {
				return err
			}
			result, err := svc.Suggest(cmd.Context(), joinArgs(args), max, raw)
			if err != nil {
				return classifyAPIError(err, flags)
			}
			return flags.printJSON(cmd, result)
		},
	}
	cmd.Flags().BoolVar(&raw, "raw", false, "Return the raw suggestion payload")
	cmd.Flags().IntVar(&max, "max", 5, "Maximum number of suggestions")
	return cmd
}

func newWortenSearchCmd(flags *rootFlags) *cobra.Command {
	var raw bool
	var contexts []string
	var page int
	cmd := &cobra.Command{
		Use:   "search <query>",
		Short: "Search Worten products with explicit context filters",
		Args: func(cmd *cobra.Command, args []string) error {
			if flags.dryRun {
				return nil
			}
			if len(args) == 0 {
				return usageErr(fmt.Errorf("search requires a query"))
			}
			if len(contexts) == 0 {
				return usageErr(fmt.Errorf("search requires at least one --context value"))
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if dryRunOK(flags) {
				return flags.printJSON(cmd, map[string]any{"dry_run": true, "command": "search"})
			}
			svc, err := newWortenService(flags)
			if err != nil {
				return err
			}
			result, err := svc.Search(cmd.Context(), joinArgs(args), contexts, page, raw)
			if err != nil {
				return classifyAPIError(err, flags)
			}
			return flags.printJSON(cmd, result)
		},
	}
	cmd.Flags().BoolVar(&raw, "raw", false, "Return the raw search payload")
	cmd.Flags().StringArrayVar(&contexts, "context", nil, "Search context value; repeat the flag for multiple contexts")
	cmd.Flags().IntVar(&page, "page", 1, "Result page number")
	return cmd
}

func newWortenSnapshotCmd(flags *rootFlags) *cobra.Command {
	var raw bool
	var refresh bool
	var cacheOnly bool
	cmd := oneInputWortenCmd(flags, "snapshot", "Capture or read a normalized Worten snapshot", func(cmd *cobra.Command, input string) (any, error) {
		svc, err := newWortenService(flags)
		if err != nil {
			return nil, err
		}
		return svc.Snapshot(cmd.Context(), input, worten.SnapshotOptions{
			Refresh:   refresh,
			CacheOnly: cacheOnly,
		}, raw)
	})
	cmd.Flags().BoolVar(&raw, "raw", false, "Return the full snapshot payload")
	cmd.Flags().BoolVar(&refresh, "refresh", false, "Force a live refresh before returning the snapshot")
	cmd.Flags().BoolVar(&cacheOnly, "cache-only", false, "Return only cached snapshots; do not hit the network")
	return cmd
}

func oneInputWortenCmd(flags *rootFlags, use, short string, runner func(*cobra.Command, string) (any, error)) *cobra.Command {
	return &cobra.Command{
		Use:   use + " <product-url-or-id>",
		Short: short,
		Args: func(cmd *cobra.Command, args []string) error {
			if flags.dryRun {
				return nil
			}
			if len(args) != 1 {
				return usageErr(fmt.Errorf("%s requires a product URL or product UUID", use))
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if dryRunOK(flags) {
				return flags.printJSON(cmd, map[string]any{"dry_run": true, "command": use})
			}
			result, err := runner(cmd, args[0])
			if err != nil {
				return classifyAPIError(err, flags)
			}
			return flags.printJSON(cmd, result)
		},
	}
}

func joinArgs(args []string) string {
	return strings.TrimSpace(strings.Join(args, " "))
}

func newWortenService(flags *rootFlags) (*worten.Service, error) {
	client, err := flags.newClient()
	if err != nil {
		return nil, err
	}
	return worten.New(client)
}
