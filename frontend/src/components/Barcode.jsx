import React, { useEffect, useRef } from "react";
import JsBarcode from "jsbarcode";

export default function Barcode({ sku }) {
  const svgRef = useRef(null);

  useEffect(() => {
    if (sku && svgRef.current) {
      JsBarcode(svgRef.current, sku, {
        format: "CODE128",
        displayValue: true,
        fontSize: 14,
        width: 2,
        height: 50,
      });
    }
  }, [sku]);

  return <svg ref={svgRef}></svg>;
}
